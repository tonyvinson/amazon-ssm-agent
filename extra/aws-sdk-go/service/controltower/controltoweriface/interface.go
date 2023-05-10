// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package controltoweriface provides an interface to enable mocking the AWS Control Tower service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package controltoweriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/controltower"
)

// ControlTowerAPI provides an interface to enable mocking the
// controltower.ControlTower service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Control Tower.
//    func myFunc(svc controltoweriface.ControlTowerAPI) bool {
//        // Make svc.DisableControl request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := controltower.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockControlTowerClient struct {
//        controltoweriface.ControlTowerAPI
//    }
//    func (m *mockControlTowerClient) DisableControl(input *controltower.DisableControlInput) (*controltower.DisableControlOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockControlTowerClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ControlTowerAPI interface {
	DisableControl(*controltower.DisableControlInput) (*controltower.DisableControlOutput, error)
	DisableControlWithContext(aws.Context, *controltower.DisableControlInput, ...request.Option) (*controltower.DisableControlOutput, error)
	DisableControlRequest(*controltower.DisableControlInput) (*request.Request, *controltower.DisableControlOutput)

	EnableControl(*controltower.EnableControlInput) (*controltower.EnableControlOutput, error)
	EnableControlWithContext(aws.Context, *controltower.EnableControlInput, ...request.Option) (*controltower.EnableControlOutput, error)
	EnableControlRequest(*controltower.EnableControlInput) (*request.Request, *controltower.EnableControlOutput)

	GetControlOperation(*controltower.GetControlOperationInput) (*controltower.GetControlOperationOutput, error)
	GetControlOperationWithContext(aws.Context, *controltower.GetControlOperationInput, ...request.Option) (*controltower.GetControlOperationOutput, error)
	GetControlOperationRequest(*controltower.GetControlOperationInput) (*request.Request, *controltower.GetControlOperationOutput)

	ListEnabledControls(*controltower.ListEnabledControlsInput) (*controltower.ListEnabledControlsOutput, error)
	ListEnabledControlsWithContext(aws.Context, *controltower.ListEnabledControlsInput, ...request.Option) (*controltower.ListEnabledControlsOutput, error)
	ListEnabledControlsRequest(*controltower.ListEnabledControlsInput) (*request.Request, *controltower.ListEnabledControlsOutput)

	ListEnabledControlsPages(*controltower.ListEnabledControlsInput, func(*controltower.ListEnabledControlsOutput, bool) bool) error
	ListEnabledControlsPagesWithContext(aws.Context, *controltower.ListEnabledControlsInput, func(*controltower.ListEnabledControlsOutput, bool) bool, ...request.Option) error
}

var _ ControlTowerAPI = (*controltower.ControlTower)(nil)
