// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	common "github.com/aws/amazon-ssm-agent/agent/update/tester/common"
	mock "github.com/stretchr/testify/mock"
)

// ITestCase is an autogenerated mock type for the ITestCase type
type ITestCase struct {
	mock.Mock
}

// CleanupTestCase provides a mock function with given fields:
func (_m *ITestCase) CleanupTestCase() {
	_m.Called()
}

// ExecuteTestCase provides a mock function with given fields:
func (_m *ITestCase) ExecuteTestCase() common.TestOutput {
	ret := _m.Called()

	var r0 common.TestOutput
	if rf, ok := ret.Get(0).(func() common.TestOutput); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.TestOutput)
	}

	return r0
}

// GetTestCaseName provides a mock function with given fields:
func (_m *ITestCase) GetTestCaseName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Initialize provides a mock function with given fields:
func (_m *ITestCase) Initialize() {
	_m.Called()
}

// ShouldRunTest provides a mock function with given fields:
func (_m *ITestCase) ShouldRunTest() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
