// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	credentialproviders "github.com/aws/amazon-ssm-agent/common/identity/credentialproviders"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ICredentialRefresherAgentIdentity is an autogenerated mock type for the ICredentialRefresherAgentIdentity type
type ICredentialRefresherAgentIdentity struct {
	mock.Mock
}

// CredentialProvider provides a mock function with given fields:
func (_m *ICredentialRefresherAgentIdentity) CredentialProvider() credentialproviders.IRemoteProvider {
	ret := _m.Called()

	var r0 credentialproviders.IRemoteProvider
	if rf, ok := ret.Get(0).(func() credentialproviders.IRemoteProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(credentialproviders.IRemoteProvider)
		}
	}

	return r0
}

// NewICredentialRefresherAgentIdentity creates a new instance of ICredentialRefresherAgentIdentity. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewICredentialRefresherAgentIdentity(t testing.TB) *ICredentialRefresherAgentIdentity {
	mock := &ICredentialRefresherAgentIdentity{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
