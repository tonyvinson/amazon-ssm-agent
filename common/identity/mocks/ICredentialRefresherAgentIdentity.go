// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	credentials "github.com/aws/aws-sdk-go/aws/credentials"

	mock "github.com/stretchr/testify/mock"
)

// ICredentialRefresherAgentIdentity is an autogenerated mock type for the ICredentialRefresherAgentIdentity type
type ICredentialRefresherAgentIdentity struct {
	mock.Mock
}

// CredentialProvider provides a mock function with given fields:
func (_m *ICredentialRefresherAgentIdentity) CredentialProvider() credentials.Provider {
	ret := _m.Called()

	var r0 credentials.Provider
	if rf, ok := ret.Get(0).(func() credentials.Provider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(credentials.Provider)
		}
	}

	return r0
}

// ShareFile provides a mock function with given fields:
func (_m *ICredentialRefresherAgentIdentity) ShareFile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ShareProfile provides a mock function with given fields:
func (_m *ICredentialRefresherAgentIdentity) ShareProfile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ShouldShareCredentials provides a mock function with given fields:
func (_m *ICredentialRefresherAgentIdentity) ShouldShareCredentials() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
