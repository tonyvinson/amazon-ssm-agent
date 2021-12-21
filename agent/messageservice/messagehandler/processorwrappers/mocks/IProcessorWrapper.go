// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contracts "github.com/aws/amazon-ssm-agent/agent/contracts"
	mock "github.com/stretchr/testify/mock"

	processor "github.com/aws/amazon-ssm-agent/agent/framework/processor"

	utils "github.com/aws/amazon-ssm-agent/agent/messageservice/utils"
)

// IProcessorWrapper is an autogenerated mock type for the IProcessorWrapper type
type IProcessorWrapper struct {
	mock.Mock
}

// GetName provides a mock function with given fields:
func (_m *IProcessorWrapper) GetName() utils.ProcessorName {
	ret := _m.Called()

	var r0 utils.ProcessorName
	if rf, ok := ret.Get(0).(func() utils.ProcessorName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(utils.ProcessorName)
	}

	return r0
}

// GetStartWorker provides a mock function with given fields:
func (_m *IProcessorWrapper) GetStartWorker() contracts.DocumentType {
	ret := _m.Called()

	var r0 contracts.DocumentType
	if rf, ok := ret.Get(0).(func() contracts.DocumentType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(contracts.DocumentType)
	}

	return r0
}

// GetTerminateWorker provides a mock function with given fields:
func (_m *IProcessorWrapper) GetTerminateWorker() contracts.DocumentType {
	ret := _m.Called()

	var r0 contracts.DocumentType
	if rf, ok := ret.Get(0).(func() contracts.DocumentType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(contracts.DocumentType)
	}

	return r0
}

// Initialize provides a mock function with given fields: outputChan
func (_m *IProcessorWrapper) Initialize(outputChan map[contracts.UpstreamServiceName]chan contracts.DocumentResult) error {
	ret := _m.Called(outputChan)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[contracts.UpstreamServiceName]chan contracts.DocumentResult) error); ok {
		r0 = rf(outputChan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushToProcessor provides a mock function with given fields: _a0
func (_m *IProcessorWrapper) PushToProcessor(_a0 contracts.DocumentState) processor.ErrorCode {
	ret := _m.Called(_a0)

	var r0 processor.ErrorCode
	if rf, ok := ret.Get(0).(func(contracts.DocumentState) processor.ErrorCode); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(processor.ErrorCode)
	}

	return r0
}

// Stop provides a mock function with given fields: stopType
func (_m *IProcessorWrapper) Stop() {
	_m.Called()
}
