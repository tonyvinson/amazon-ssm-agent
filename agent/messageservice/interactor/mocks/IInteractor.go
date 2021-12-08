// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	utils "github.com/aws/amazon-ssm-agent/agent/messageservice/utils"
	mock "github.com/stretchr/testify/mock"
)

// IInteractor is an autogenerated mock type for the IInteractor type
type IInteractor struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *IInteractor) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *IInteractor) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSupportedWorkers provides a mock function with given fields:
func (_m *IInteractor) GetSupportedWorkers() []utils.WorkerName {
	ret := _m.Called()

	var r0 []utils.WorkerName
	if rf, ok := ret.Get(0).(func() []utils.WorkerName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]utils.WorkerName)
		}
	}

	return r0
}

// Initialize provides a mock function with given fields:
func (_m *IInteractor) Initialize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostProcessorInitialization provides a mock function with given fields:
func (_m *IInteractor) PostProcessorInitialization() {
	_m.Called()
}

// PreProcessorClose provides a mock function with given fields:
func (_m *IInteractor) PreProcessorClose() {
	_m.Called()
}
