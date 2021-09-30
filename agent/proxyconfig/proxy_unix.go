//go:build freebsd || linux || netbsd || openbsd
// +build freebsd linux netbsd openbsd

// Copyright 2021 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Package proxy config to handle set/get proxy settings
package proxyconfig

import "github.com/aws/amazon-ssm-agent/agent/log"

// SetProxyConfig on linux returns the environment variables right away
// because using environment variables is the only way to configure proxy on linux
func SetProxyConfig(log log.T) map[string]string {
	return GetProxyConfig()
}
