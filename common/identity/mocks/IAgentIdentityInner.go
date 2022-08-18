// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	credentials "github.com/aws/aws-sdk-go/aws/credentials"
	mock "github.com/stretchr/testify/mock"
)

// IAgentIdentityInner is an autogenerated mock type for the IAgentIdentityInner type
type IAgentIdentityInner struct {
	mock.Mock
}

// AvailabilityZone provides a mock function with given fields:
func (_m *IAgentIdentityInner) AvailabilityZone() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AvailabilityZoneId provides a mock function with given fields:
func (_m *IAgentIdentityInner) AvailabilityZoneId() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Credentials provides a mock function with given fields:
func (_m *IAgentIdentityInner) Credentials() *credentials.Credentials {
	ret := _m.Called()

	var r0 *credentials.Credentials
	if rf, ok := ret.Get(0).(func() *credentials.Credentials); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*credentials.Credentials)
		}
	}

	return r0
}

// IdentityType provides a mock function with given fields:
func (_m *IAgentIdentityInner) IdentityType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InstanceID provides a mock function with given fields:
func (_m *IAgentIdentityInner) InstanceID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstanceType provides a mock function with given fields:
func (_m *IAgentIdentityInner) InstanceType() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsIdentityEnvironment provides a mock function with given fields:
func (_m *IAgentIdentityInner) IsIdentityEnvironment() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Region provides a mock function with given fields:
func (_m *IAgentIdentityInner) Region() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceDomain provides a mock function with given fields:
func (_m *IAgentIdentityInner) ServiceDomain() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VpcPrimaryCIDRBlock provides a mock function with given fields:
func (_m *IAgentIdentityInner) VpcPrimaryCIDRBlock() (map[string][]string, error) {
	ret := _m.Called()

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
