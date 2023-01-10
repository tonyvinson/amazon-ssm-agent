/*
 * Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"). You may not
 * use this file except in compliance with the License. A copy of the
 * License is located at
 *
 * http://aws.amazon.com/apache2.0/
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

// Package mock defines the struct and its corresponding methods for mocking core.Repository
package mock

import (
	"github.com/aws/amazon-ssm-agent/agent/plugins/downloadcontent/gitresource/privategit/handler/core"
	"github.com/stretchr/testify/mock"
)

type GitRepositoryMock struct {
	mock.Mock
}

func (gitRepository *GitRepositoryMock) Worktree() (gitWorktree core.IGitWorktree, err error) {
	args := gitRepository.Called()
	return args.Get(0).(core.IGitWorktree), args.Error(1)
}
