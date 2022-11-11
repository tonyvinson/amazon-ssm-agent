// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	log "github.com/aws/amazon-ssm-agent/agent/log"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// IOnpremRegistrationInfo is an autogenerated mock type for the IOnpremRegistrationInfo type
type IOnpremRegistrationInfo struct {
	mock.Mock
}

// Fingerprint provides a mock function with given fields: _a0
func (_m *IOnpremRegistrationInfo) Fingerprint(_a0 log.T) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(log.T) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(log.T) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateKeyPair provides a mock function with given fields:
func (_m *IOnpremRegistrationInfo) GenerateKeyPair() (string, string, string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func() string); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GeneratePublicKey provides a mock function with given fields: _a0
func (_m *IOnpremRegistrationInfo) GeneratePublicKey(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasManagedInstancesCredentials provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) HasManagedInstancesCredentials(_a0 log.T, _a1 string, _a2 string) bool {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(log.T, string, string) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InstanceID provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) InstanceID(_a0 log.T, _a1 string, _a2 string) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(log.T, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrivateKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) PrivateKey(_a0 log.T, _a1 string, _a2 string) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(log.T, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrivateKeyType provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) PrivateKeyType(_a0 log.T, _a1 string, _a2 string) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(log.T, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Region provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) Region(_a0 log.T, _a1 string, _a2 string) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(log.T, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReloadInstanceInfo provides a mock function with given fields: _a0, _a1, _a2
func (_m *IOnpremRegistrationInfo) ReloadInstanceInfo(_a0 log.T, _a1 string, _a2 string) {
	_m.Called(_a0, _a1, _a2)
}

// ShouldRotatePrivateKey provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *IOnpremRegistrationInfo) ShouldRotatePrivateKey(_a0 log.T, _a1 string, _a2 int, _a3 bool, _a4 string, _a5 string) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 bool
	if rf, ok := ret.Get(0).(func(log.T, string, int, bool, string, string) bool); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(log.T, string, int, bool, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePrivateKey provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *IOnpremRegistrationInfo) UpdatePrivateKey(_a0 log.T, _a1 string, _a2 string, _a3 string, _a4 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T, string, string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIOnpremRegistrationInfo creates a new instance of IOnpremRegistrationInfo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIOnpremRegistrationInfo(t testing.TB) *IOnpremRegistrationInfo {
	mock := &IOnpremRegistrationInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
