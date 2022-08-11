// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package onprem

import (
	"sync"

	"github.com/aws/amazon-ssm-agent/common/identity/credentialproviders"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/aws/amazon-ssm-agent/agent/managedInstances/registration"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	// IdentityType is the identity type for OnPrem
	IdentityType = "OnPrem"
)

// Identity is the struct defining the IAgentIdentityInner for OnPrem
type Identity struct {
	Log                    log.T
	Config                 *appconfig.SsmagentConfig
	registrationInfo       registration.IOnpremRegistrationInfo
	credentials            *credentials.Credentials
	credentialsProvider    credentialproviders.IRemoteProvider
	shareFile              string
	shouldShareCredentials bool
	credsInitMutex         sync.Mutex
}
