// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package runcommand implements runcommand core processing module
package runcommand

import (
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/fileutil"
	"github.com/aws/amazon-ssm-agent/agent/framework/docmanager"
	messageContracts "github.com/aws/amazon-ssm-agent/agent/runcommand/contracts"
	mdsService "github.com/aws/amazon-ssm-agent/agent/runcommand/mds"
	"github.com/aws/amazon-ssm-agent/agent/sdkutil"
	"github.com/aws/aws-sdk-go/service/ssmmds"
	"github.com/carlescere/scheduler"
)

const (
	documentContent  = "DocumentContent"
	runtimeConfig    = "runtimeConfig"
	cloudwatchPlugin = "aws:cloudWatch"
	properties       = "properties"
	parameters       = "Parameters"
	// MDS service will mark document as timeout if it didn't recieve any responce from the agent after 2 hours
	documentLevelTimeOutDurationHour = 2
)

var singletonMapOfUnsupportedSSMDocs map[string]bool
var once sync.Once

var loadDocStateFromSendCommand = parseSendCommandMessage
var loadDocStateFromCancelCommand = parseCancelCommandMessage

// Name returns the module name
func (s *RunCommandService) ModuleName() string {
	return s.name
}

// Execute starts the scheduling of the message processor plugin
func (s *RunCommandService) ModuleExecute() (err error) {
	log := s.context.Log()
	defer func() {
		if msg := recover(); msg != nil {
			log.Errorf("run command ModuleExecute run panic: %v", msg)
		}
	}()
	log.Info("Starting document processing engine...")
	var resultChan chan contracts.DocumentResult
	if resultChan, err = s.processor.Start(); err != nil {
		log.Errorf("unable to start document processor: %v", err)
		return
	}

	go s.listenReply(resultChan)

	if err = s.processor.InitialProcessing(true); err != nil {
		log.Errorf("initial processing in EngineProcessor encountered error: %v", err)
		return
	}

	log.Info("Scheduling message polling")
	s.messagePollWaitGroup = &sync.WaitGroup{}
	if s.messagePollJob, err = scheduler.Every(pollMessageFrequencyMinutes).Minutes().Run(s.messagePollLoop); err != nil {
		s.context.Log().Errorf("unable to schedule message poll job. %v", err)
	}

	log.Info("Starting send replies to MDS")
	if s.sendReplyJob, err = scheduler.Every(sendReplyFrequencyMinutes).Minutes().Run(s.sendReplyLoop); err != nil {
		s.context.Log().Errorf("unable to schedule send reply job. %v", err)
	}

	//TODO move association polling out in the next CR
	if s.pollAssociations {
		s.assocProcessor.ModuleExecute()
	}
	return
}

func (s *RunCommandService) ModuleRequestStop(stopType contracts.StopType) (err error) {
	//first stop sending failed replies to the service and the message poller
	s.stop()
	//second stop the message processor
	s.processor.Stop(stopType)

	//TODO move this out once we have association moved to a different core module
	if s.assocProcessor != nil {
		s.assocProcessor.ModuleRequestStop(stopType)
	}
	return nil
}

func (s *RunCommandService) listenReply(resultChan chan contracts.DocumentResult) {
	log := s.context.Log()
	//processor guarantees to close this channel upon stop
	for res := range resultChan {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("Failed to process replies with error %v", err)
					log.Errorf("Stracktrace:\n%s", debug.Stack())
				}
			}()

			//cloudwatch and refresh association needs to trigger the in-memory component, adding filter here
			s.handleSpecialPlugin(res.LastPlugin, res.PluginResults, res.MessageID)

			if res.LastPlugin != "" {
				log.Infof("received plugin: %v result from Processor", res.LastPlugin)
			} else {
				log.Infof("command: %v complete", res.MessageID)
				//Deleting Old Log Files after the execution is over and files have been moved to completed folder
				//clean completed document state files and orchestration dirs. Takes care of only files generated by RunCommand in the folder
				shortInstanceId, _ := s.context.Identity().ShortInstanceID()
				go docmanager.DeleteOldOrchestrationDirectories(log,
					shortInstanceId,
					s.context.AppConfig().Agent.OrchestrationRootDir,
					s.context.AppConfig().Ssm.RunCommandLogsRetentionDurationHours,
					s.context.AppConfig().Ssm.AssociationLogsRetentionDurationHours)
			}
			s.sendResponse(res.MessageID, res)
		}()
	}
}

//temporary solution on plugins with shared responsibility with agent
func (s *RunCommandService) handleSpecialPlugin(lastPluginID string, pluginRes map[string]*contracts.PluginResult, messageID string) {
	var newRes contracts.PluginResult

	log := s.context.Log()
	//TODO once association service switch to use RC and CW goes away, remove this block
	for ID, pluginRes := range pluginRes {
		if pluginRes.PluginName == appconfig.PluginNameRefreshAssociation {
			log.Infof("Found %v to invoke refresh association immediately", pluginRes.PluginName)
			commandID, _ := messageContracts.GetCommandID(messageID)
			orchestrationDir := fileutil.BuildPath(s.orchestrationRootDir, commandID)
			//apply association only when this is the last plugin run
			s.assocProcessor.ProcessRefreshAssociation(log, pluginRes, orchestrationDir, lastPluginID == ID)

			log.Infof("Finished refreshing association immediately - response: %v", newRes)
		}
	}

}

func (s *RunCommandService) processMessage(msg *ssmmds.Message) {
	var (
		docState *contracts.DocumentState
		err      error
	)

	// create separate logger that includes messageID with every log message
	context := s.context.With("[messageID=" + *msg.MessageId + "]")
	log := context.Log()
	log.Debug("Processing message")

	if err = validate(msg); err != nil {
		log.Error("message not valid, ignoring: ", err)
		return
	}

	if strings.HasPrefix(*msg.Topic, string(SendCommandTopicPrefix)) {
		docState, err = loadDocStateFromSendCommand(context, msg, s.orchestrationRootDir)
		if err != nil {
			log.Error(err)
			s.sendDocLevelResponse(*msg.MessageId, contracts.ResultStatusFailed, err.Error())
			return
		}
	} else if strings.HasPrefix(*msg.Topic, string(CancelCommandTopicPrefix)) {
		docState, err = loadDocStateFromCancelCommand(context, msg, s.orchestrationRootDir)
	} else {
		err = fmt.Errorf("unexpected topic name %v", *msg.Topic)
	}

	if err != nil {
		log.Error("format of received message is invalid ", err)
		if err = s.service.FailMessage(log, *msg.MessageId, mdsService.InternalHandlerException); err != nil {
			sdkutil.HandleAwsError(log, err, s.processorStopPolicy)
		}
		return
	}
	if err = s.service.AcknowledgeMessage(log, *msg.MessageId); err != nil {
		sdkutil.HandleAwsError(log, err, s.processorStopPolicy)
		return
	}

	log.Debugf("Ack done. Received message - messageId - %v", *msg.MessageId)

	log.Debugf("Processing to send a reply to update the document status to InProgress")

	//TODO This function should be called in service when it submits the document to the engine
	s.sendDocLevelResponse(*msg.MessageId, contracts.ResultStatusInProgress, "")

	log.Debugf("SendReply done. Received message - messageId - %v", *msg.MessageId)
	switch docState.DocumentType {
	case contracts.SendCommand, contracts.SendCommandOffline:
		s.processor.Submit(*docState)
	case contracts.CancelCommand, contracts.CancelCommandOffline:
		s.processor.Cancel(*docState)

	default:
		log.Error("unexpected document type ", docState.DocumentType)
	}

}

// sendFailedReplies loads replies from local disk and send it again to the service, if it fails no action is needed
func (s *RunCommandService) sendFailedReplies() {
	log := s.context.Log()

	log.Debug("Checking if there are document replies that failed to reach the service, and retry sending them")
	replies := s.service.LoadFailedReplies(log)

	if len(replies) != 0 {
		log.Infof("Found document replies that need to be sent to the service")
		for _, reply := range replies {
			log.Debug("Loading reply ", reply)
			if isValidReplyRequest(reply) == false {
				log.Debug("Reply is old, document execution must have timed out. Deleting the reply")
				s.service.DeleteFailedReply(log, reply)
				continue
			}
			sendReplyRequest, err := s.service.GetFailedReply(log, reply)
			if err != nil {
				log.Error("Couldn't load the reply from disk ", err)
				continue
			}

			log.Info("Sending reply ", reply)
			if err = s.service.SendReplyWithInput(log, sendReplyRequest); err != nil {
				sdkutil.HandleAwsError(log, err, s.processorStopPolicy)
				break
			} else {
				log.Infof("Sending reply %v succeeded, deleting the reply file from disk", reply)
				s.service.DeleteFailedReply(log, reply)
			}
		}
	} else {
		log.Debugf("No failed document replies found")
	}
}

// isValidReplyRequest checks if the sendReply request is older than 2 hours
// If so it is considered as not valid anymore as the document must have timed out
func isValidReplyRequest(filename string) bool {
	splitFileName := strings.Split(filename, "_")
	if len(splitFileName) < 2 {
		return false
	}
	t, _ := time.Parse("2006-01-02T15-04-05", splitFileName[1])
	curTime := time.Now().UTC()
	delta := curTime.Sub(t).Hours()
	if delta > documentLevelTimeOutDurationHour {
		return false
	} else {
		return true
	}
}
