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

// Package processor contains the methods for update ssm agent.
// It also provides methods for sendReply and updateInstanceInfo
package processor

import (
	"testing"
	"time"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/log"
	contextmocks "github.com/aws/amazon-ssm-agent/agent/mocks/context"
	logmocks "github.com/aws/amazon-ssm-agent/agent/mocks/log"
	messageService "github.com/aws/amazon-ssm-agent/agent/runcommand/mds"
	"github.com/aws/aws-sdk-go/service/ssmmds"
	"github.com/stretchr/testify/assert"
)

// stubSdkService is the stub for sdkService
type stubSdkService struct{}

func (s *stubSdkService) GetMessages(log log.T, instanceID string) (messages *ssmmds.GetMessagesOutput, err error) {
	return &ssmmds.GetMessagesOutput{}, nil
}

func (s *stubSdkService) AcknowledgeMessage(log log.T, messageID string) error {
	return nil
}

func (s *stubSdkService) SendReply(log log.T, messageID string, payload string) error {
	return nil
}

func (s *stubSdkService) FailMessage(log log.T, messageID string, failureType messageService.FailureType) error {
	return nil
}

func (s *stubSdkService) DeleteMessage(log log.T, messageID string) error {
	return nil
}

func (s *stubSdkService) Stop() {}

func (s *stubSdkService) LoadFailedReplies(log log.T) []string {
	return nil
}

func (s *stubSdkService) DeleteFailedReply(log log.T, replyId string) {}

func (s *stubSdkService) PersistFailedReply(log log.T, sendReply ssmmds.SendReplyInput) error {
	return nil
}

func (s *stubSdkService) GetFailedReply(log log.T, replyId string) (*ssmmds.SendReplyInput, error) {
	return nil, nil
}

func (s *stubSdkService) SendReplyWithInput(log log.T, sendReply *ssmmds.SendReplyInput) error {
	return nil
}

func stubNewMsgSvc(context context.T, connectionTimeout time.Duration) messageService.Service {
	return &stubSdkService{}
}

func TestSendReply(t *testing.T) {
	var logger = logmocks.NewMockLog()
	updateDetail := createUpdateDetail(Installed)
	service := svcManager{
		context: contextmocks.NewMockDefault(),
	}
	// setup
	getAppConfig = func(bool) (appconfig.SsmagentConfig, error) {
		config := appconfig.SsmagentConfig{}
		return config, nil
	}

	newMsgSvc = stubNewMsgSvc

	// action
	err := service.SendReply(logger, updateDetail)

	// assert
	assert.NoError(t, err)
}

func TestSendReplyDeleteMessage(t *testing.T) {
	var logger = logmocks.NewMockLog()
	updateDetail := createUpdateDetail(Installed)
	service := svcManager{
		context: contextmocks.NewMockDefault(),
	}
	// setup
	getAppConfig = func(bool) (appconfig.SsmagentConfig, error) {
		config := appconfig.SsmagentConfig{}
		return config, nil
	}
	newMsgSvc = stubNewMsgSvc

	// action
	err := service.DeleteMessage(logger, updateDetail)

	// assert
	assert.NoError(t, err)
}

func TestPrepareAgentResultCorrectlyPopulatesFields(t *testing.T) {
	context := contextmocks.NewMockDefault()
	updateDetail := createUpdateDetail(Installed)

	agentResult := prepareAgentResult(context, updateDetail)
	assert.Equal(t, updateDetail.MessageID, agentResult.MessageID)
	assert.Equal(t, updateDetail.Result, agentResult.Status)
	assert.Equal(t, 1, agentResult.NPlugins)
	assert.Equal(t, contracts.RunCommandResult, agentResult.ResultType)
	assert.Equal(t, contracts.SendCommand, agentResult.RelatedDocumentType)
}

func TestPreparePluginResultsPopulatesCodeOneForFailedUpdate(t *testing.T) {
	mockContext := contextmocks.NewMockDefault()
	updateDetail := &UpdateDetail{
		Result: contracts.ResultStatusFailed,
	}

	pluginResults := preparePluginResults(mockContext, updateDetail)
	assert.Equal(t, 1, pluginResults[appconfig.PluginNameAwsAgentUpdate].Code)
}

func TestPreparePluginResultsCorrectlyPopulatesPlugin(t *testing.T) {
	mockContext := contextmocks.NewMockDefault()
	startTime := time.Now()
	endTime := time.Now()

	updateDetail := &UpdateDetail{
		Result:             contracts.ResultStatusSuccess,
		StartDateTime:      startTime,
		EndDateTime:        endTime,
		StandardOut:        "out",
		StandardError:      "err",
		OutputS3BucketName: "bucketName",
		OutputS3KeyPrefix:  "keyPrefix",
	}

	pluginResults := preparePluginResults(mockContext, updateDetail)

	expectedPluginResult := &contracts.PluginResult{
		Status:             contracts.ResultStatusSuccess,
		Code:               0,
		Output:             "out\n----------ERROR-------\nerr",
		StartDateTime:      startTime,
		EndDateTime:        pluginResults[appconfig.PluginNameAwsAgentUpdate].EndDateTime,
		OutputS3BucketName: "bucketName",
		OutputS3KeyPrefix:  "keyPrefix",
		StandardOutput:     "out",
		StandardError:      "err",
	}

	assert.True(t, len(pluginResults) == 1)
	assert.Equal(t, expectedPluginResult, pluginResults[appconfig.PluginNameAwsAgentUpdate])
}
