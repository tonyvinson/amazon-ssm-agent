// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package s3controliface provides an interface to enable mocking the AWS S3 Control service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package s3controliface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3control"
)

// S3ControlAPI provides an interface to enable mocking the
// s3control.S3Control service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS S3 Control.
//    func myFunc(svc s3controliface.S3ControlAPI) bool {
//        // Make svc.CreateAccessPoint request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := s3control.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockS3ControlClient struct {
//        s3controliface.S3ControlAPI
//    }
//    func (m *mockS3ControlClient) CreateAccessPoint(input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockS3ControlClient{}
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
type S3ControlAPI interface {
	CreateAccessPoint(*s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointWithContext(aws.Context, *s3control.CreateAccessPointInput, ...request.Option) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointRequest(*s3control.CreateAccessPointInput) (*request.Request, *s3control.CreateAccessPointOutput)

	CreateAccessPointForObjectLambda(*s3control.CreateAccessPointForObjectLambdaInput) (*s3control.CreateAccessPointForObjectLambdaOutput, error)
	CreateAccessPointForObjectLambdaWithContext(aws.Context, *s3control.CreateAccessPointForObjectLambdaInput, ...request.Option) (*s3control.CreateAccessPointForObjectLambdaOutput, error)
	CreateAccessPointForObjectLambdaRequest(*s3control.CreateAccessPointForObjectLambdaInput) (*request.Request, *s3control.CreateAccessPointForObjectLambdaOutput)

	CreateBucket(*s3control.CreateBucketInput) (*s3control.CreateBucketOutput, error)
	CreateBucketWithContext(aws.Context, *s3control.CreateBucketInput, ...request.Option) (*s3control.CreateBucketOutput, error)
	CreateBucketRequest(*s3control.CreateBucketInput) (*request.Request, *s3control.CreateBucketOutput)

	CreateJob(*s3control.CreateJobInput) (*s3control.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *s3control.CreateJobInput, ...request.Option) (*s3control.CreateJobOutput, error)
	CreateJobRequest(*s3control.CreateJobInput) (*request.Request, *s3control.CreateJobOutput)

	CreateMultiRegionAccessPoint(*s3control.CreateMultiRegionAccessPointInput) (*s3control.CreateMultiRegionAccessPointOutput, error)
	CreateMultiRegionAccessPointWithContext(aws.Context, *s3control.CreateMultiRegionAccessPointInput, ...request.Option) (*s3control.CreateMultiRegionAccessPointOutput, error)
	CreateMultiRegionAccessPointRequest(*s3control.CreateMultiRegionAccessPointInput) (*request.Request, *s3control.CreateMultiRegionAccessPointOutput)

	DeleteAccessPoint(*s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointWithContext(aws.Context, *s3control.DeleteAccessPointInput, ...request.Option) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointRequest(*s3control.DeleteAccessPointInput) (*request.Request, *s3control.DeleteAccessPointOutput)

	DeleteAccessPointForObjectLambda(*s3control.DeleteAccessPointForObjectLambdaInput) (*s3control.DeleteAccessPointForObjectLambdaOutput, error)
	DeleteAccessPointForObjectLambdaWithContext(aws.Context, *s3control.DeleteAccessPointForObjectLambdaInput, ...request.Option) (*s3control.DeleteAccessPointForObjectLambdaOutput, error)
	DeleteAccessPointForObjectLambdaRequest(*s3control.DeleteAccessPointForObjectLambdaInput) (*request.Request, *s3control.DeleteAccessPointForObjectLambdaOutput)

	DeleteAccessPointPolicy(*s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyWithContext(aws.Context, *s3control.DeleteAccessPointPolicyInput, ...request.Option) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyRequest(*s3control.DeleteAccessPointPolicyInput) (*request.Request, *s3control.DeleteAccessPointPolicyOutput)

	DeleteAccessPointPolicyForObjectLambda(*s3control.DeleteAccessPointPolicyForObjectLambdaInput) (*s3control.DeleteAccessPointPolicyForObjectLambdaOutput, error)
	DeleteAccessPointPolicyForObjectLambdaWithContext(aws.Context, *s3control.DeleteAccessPointPolicyForObjectLambdaInput, ...request.Option) (*s3control.DeleteAccessPointPolicyForObjectLambdaOutput, error)
	DeleteAccessPointPolicyForObjectLambdaRequest(*s3control.DeleteAccessPointPolicyForObjectLambdaInput) (*request.Request, *s3control.DeleteAccessPointPolicyForObjectLambdaOutput)

	DeleteBucket(*s3control.DeleteBucketInput) (*s3control.DeleteBucketOutput, error)
	DeleteBucketWithContext(aws.Context, *s3control.DeleteBucketInput, ...request.Option) (*s3control.DeleteBucketOutput, error)
	DeleteBucketRequest(*s3control.DeleteBucketInput) (*request.Request, *s3control.DeleteBucketOutput)

	DeleteBucketLifecycleConfiguration(*s3control.DeleteBucketLifecycleConfigurationInput) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketLifecycleConfigurationWithContext(aws.Context, *s3control.DeleteBucketLifecycleConfigurationInput, ...request.Option) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketLifecycleConfigurationRequest(*s3control.DeleteBucketLifecycleConfigurationInput) (*request.Request, *s3control.DeleteBucketLifecycleConfigurationOutput)

	DeleteBucketPolicy(*s3control.DeleteBucketPolicyInput) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyWithContext(aws.Context, *s3control.DeleteBucketPolicyInput, ...request.Option) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyRequest(*s3control.DeleteBucketPolicyInput) (*request.Request, *s3control.DeleteBucketPolicyOutput)

	DeleteBucketReplication(*s3control.DeleteBucketReplicationInput) (*s3control.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationWithContext(aws.Context, *s3control.DeleteBucketReplicationInput, ...request.Option) (*s3control.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationRequest(*s3control.DeleteBucketReplicationInput) (*request.Request, *s3control.DeleteBucketReplicationOutput)

	DeleteBucketTagging(*s3control.DeleteBucketTaggingInput) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingWithContext(aws.Context, *s3control.DeleteBucketTaggingInput, ...request.Option) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingRequest(*s3control.DeleteBucketTaggingInput) (*request.Request, *s3control.DeleteBucketTaggingOutput)

	DeleteJobTagging(*s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error)
	DeleteJobTaggingWithContext(aws.Context, *s3control.DeleteJobTaggingInput, ...request.Option) (*s3control.DeleteJobTaggingOutput, error)
	DeleteJobTaggingRequest(*s3control.DeleteJobTaggingInput) (*request.Request, *s3control.DeleteJobTaggingOutput)

	DeleteMultiRegionAccessPoint(*s3control.DeleteMultiRegionAccessPointInput) (*s3control.DeleteMultiRegionAccessPointOutput, error)
	DeleteMultiRegionAccessPointWithContext(aws.Context, *s3control.DeleteMultiRegionAccessPointInput, ...request.Option) (*s3control.DeleteMultiRegionAccessPointOutput, error)
	DeleteMultiRegionAccessPointRequest(*s3control.DeleteMultiRegionAccessPointInput) (*request.Request, *s3control.DeleteMultiRegionAccessPointOutput)

	DeletePublicAccessBlock(*s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockWithContext(aws.Context, *s3control.DeletePublicAccessBlockInput, ...request.Option) (*s3control.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockRequest(*s3control.DeletePublicAccessBlockInput) (*request.Request, *s3control.DeletePublicAccessBlockOutput)

	DeleteStorageLensConfiguration(*s3control.DeleteStorageLensConfigurationInput) (*s3control.DeleteStorageLensConfigurationOutput, error)
	DeleteStorageLensConfigurationWithContext(aws.Context, *s3control.DeleteStorageLensConfigurationInput, ...request.Option) (*s3control.DeleteStorageLensConfigurationOutput, error)
	DeleteStorageLensConfigurationRequest(*s3control.DeleteStorageLensConfigurationInput) (*request.Request, *s3control.DeleteStorageLensConfigurationOutput)

	DeleteStorageLensConfigurationTagging(*s3control.DeleteStorageLensConfigurationTaggingInput) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	DeleteStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.DeleteStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	DeleteStorageLensConfigurationTaggingRequest(*s3control.DeleteStorageLensConfigurationTaggingInput) (*request.Request, *s3control.DeleteStorageLensConfigurationTaggingOutput)

	DescribeJob(*s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error)
	DescribeJobWithContext(aws.Context, *s3control.DescribeJobInput, ...request.Option) (*s3control.DescribeJobOutput, error)
	DescribeJobRequest(*s3control.DescribeJobInput) (*request.Request, *s3control.DescribeJobOutput)

	DescribeMultiRegionAccessPointOperation(*s3control.DescribeMultiRegionAccessPointOperationInput) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error)
	DescribeMultiRegionAccessPointOperationWithContext(aws.Context, *s3control.DescribeMultiRegionAccessPointOperationInput, ...request.Option) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error)
	DescribeMultiRegionAccessPointOperationRequest(*s3control.DescribeMultiRegionAccessPointOperationInput) (*request.Request, *s3control.DescribeMultiRegionAccessPointOperationOutput)

	GetAccessPoint(*s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error)
	GetAccessPointWithContext(aws.Context, *s3control.GetAccessPointInput, ...request.Option) (*s3control.GetAccessPointOutput, error)
	GetAccessPointRequest(*s3control.GetAccessPointInput) (*request.Request, *s3control.GetAccessPointOutput)

	GetAccessPointConfigurationForObjectLambda(*s3control.GetAccessPointConfigurationForObjectLambdaInput) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	GetAccessPointConfigurationForObjectLambdaWithContext(aws.Context, *s3control.GetAccessPointConfigurationForObjectLambdaInput, ...request.Option) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	GetAccessPointConfigurationForObjectLambdaRequest(*s3control.GetAccessPointConfigurationForObjectLambdaInput) (*request.Request, *s3control.GetAccessPointConfigurationForObjectLambdaOutput)

	GetAccessPointForObjectLambda(*s3control.GetAccessPointForObjectLambdaInput) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	GetAccessPointForObjectLambdaWithContext(aws.Context, *s3control.GetAccessPointForObjectLambdaInput, ...request.Option) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	GetAccessPointForObjectLambdaRequest(*s3control.GetAccessPointForObjectLambdaInput) (*request.Request, *s3control.GetAccessPointForObjectLambdaOutput)

	GetAccessPointPolicy(*s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyWithContext(aws.Context, *s3control.GetAccessPointPolicyInput, ...request.Option) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyRequest(*s3control.GetAccessPointPolicyInput) (*request.Request, *s3control.GetAccessPointPolicyOutput)

	GetAccessPointPolicyForObjectLambda(*s3control.GetAccessPointPolicyForObjectLambdaInput) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	GetAccessPointPolicyForObjectLambdaWithContext(aws.Context, *s3control.GetAccessPointPolicyForObjectLambdaInput, ...request.Option) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	GetAccessPointPolicyForObjectLambdaRequest(*s3control.GetAccessPointPolicyForObjectLambdaInput) (*request.Request, *s3control.GetAccessPointPolicyForObjectLambdaOutput)

	GetAccessPointPolicyStatus(*s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusWithContext(aws.Context, *s3control.GetAccessPointPolicyStatusInput, ...request.Option) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusRequest(*s3control.GetAccessPointPolicyStatusInput) (*request.Request, *s3control.GetAccessPointPolicyStatusOutput)

	GetAccessPointPolicyStatusForObjectLambda(*s3control.GetAccessPointPolicyStatusForObjectLambdaInput) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	GetAccessPointPolicyStatusForObjectLambdaWithContext(aws.Context, *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, ...request.Option) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	GetAccessPointPolicyStatusForObjectLambdaRequest(*s3control.GetAccessPointPolicyStatusForObjectLambdaInput) (*request.Request, *s3control.GetAccessPointPolicyStatusForObjectLambdaOutput)

	GetBucket(*s3control.GetBucketInput) (*s3control.GetBucketOutput, error)
	GetBucketWithContext(aws.Context, *s3control.GetBucketInput, ...request.Option) (*s3control.GetBucketOutput, error)
	GetBucketRequest(*s3control.GetBucketInput) (*request.Request, *s3control.GetBucketOutput)

	GetBucketLifecycleConfiguration(*s3control.GetBucketLifecycleConfigurationInput) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationWithContext(aws.Context, *s3control.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationRequest(*s3control.GetBucketLifecycleConfigurationInput) (*request.Request, *s3control.GetBucketLifecycleConfigurationOutput)

	GetBucketPolicy(*s3control.GetBucketPolicyInput) (*s3control.GetBucketPolicyOutput, error)
	GetBucketPolicyWithContext(aws.Context, *s3control.GetBucketPolicyInput, ...request.Option) (*s3control.GetBucketPolicyOutput, error)
	GetBucketPolicyRequest(*s3control.GetBucketPolicyInput) (*request.Request, *s3control.GetBucketPolicyOutput)

	GetBucketReplication(*s3control.GetBucketReplicationInput) (*s3control.GetBucketReplicationOutput, error)
	GetBucketReplicationWithContext(aws.Context, *s3control.GetBucketReplicationInput, ...request.Option) (*s3control.GetBucketReplicationOutput, error)
	GetBucketReplicationRequest(*s3control.GetBucketReplicationInput) (*request.Request, *s3control.GetBucketReplicationOutput)

	GetBucketTagging(*s3control.GetBucketTaggingInput) (*s3control.GetBucketTaggingOutput, error)
	GetBucketTaggingWithContext(aws.Context, *s3control.GetBucketTaggingInput, ...request.Option) (*s3control.GetBucketTaggingOutput, error)
	GetBucketTaggingRequest(*s3control.GetBucketTaggingInput) (*request.Request, *s3control.GetBucketTaggingOutput)

	GetBucketVersioning(*s3control.GetBucketVersioningInput) (*s3control.GetBucketVersioningOutput, error)
	GetBucketVersioningWithContext(aws.Context, *s3control.GetBucketVersioningInput, ...request.Option) (*s3control.GetBucketVersioningOutput, error)
	GetBucketVersioningRequest(*s3control.GetBucketVersioningInput) (*request.Request, *s3control.GetBucketVersioningOutput)

	GetJobTagging(*s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error)
	GetJobTaggingWithContext(aws.Context, *s3control.GetJobTaggingInput, ...request.Option) (*s3control.GetJobTaggingOutput, error)
	GetJobTaggingRequest(*s3control.GetJobTaggingInput) (*request.Request, *s3control.GetJobTaggingOutput)

	GetMultiRegionAccessPoint(*s3control.GetMultiRegionAccessPointInput) (*s3control.GetMultiRegionAccessPointOutput, error)
	GetMultiRegionAccessPointWithContext(aws.Context, *s3control.GetMultiRegionAccessPointInput, ...request.Option) (*s3control.GetMultiRegionAccessPointOutput, error)
	GetMultiRegionAccessPointRequest(*s3control.GetMultiRegionAccessPointInput) (*request.Request, *s3control.GetMultiRegionAccessPointOutput)

	GetMultiRegionAccessPointPolicy(*s3control.GetMultiRegionAccessPointPolicyInput) (*s3control.GetMultiRegionAccessPointPolicyOutput, error)
	GetMultiRegionAccessPointPolicyWithContext(aws.Context, *s3control.GetMultiRegionAccessPointPolicyInput, ...request.Option) (*s3control.GetMultiRegionAccessPointPolicyOutput, error)
	GetMultiRegionAccessPointPolicyRequest(*s3control.GetMultiRegionAccessPointPolicyInput) (*request.Request, *s3control.GetMultiRegionAccessPointPolicyOutput)

	GetMultiRegionAccessPointPolicyStatus(*s3control.GetMultiRegionAccessPointPolicyStatusInput) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error)
	GetMultiRegionAccessPointPolicyStatusWithContext(aws.Context, *s3control.GetMultiRegionAccessPointPolicyStatusInput, ...request.Option) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error)
	GetMultiRegionAccessPointPolicyStatusRequest(*s3control.GetMultiRegionAccessPointPolicyStatusInput) (*request.Request, *s3control.GetMultiRegionAccessPointPolicyStatusOutput)

	GetMultiRegionAccessPointRoutes(*s3control.GetMultiRegionAccessPointRoutesInput) (*s3control.GetMultiRegionAccessPointRoutesOutput, error)
	GetMultiRegionAccessPointRoutesWithContext(aws.Context, *s3control.GetMultiRegionAccessPointRoutesInput, ...request.Option) (*s3control.GetMultiRegionAccessPointRoutesOutput, error)
	GetMultiRegionAccessPointRoutesRequest(*s3control.GetMultiRegionAccessPointRoutesInput) (*request.Request, *s3control.GetMultiRegionAccessPointRoutesOutput)

	GetPublicAccessBlock(*s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockWithContext(aws.Context, *s3control.GetPublicAccessBlockInput, ...request.Option) (*s3control.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockRequest(*s3control.GetPublicAccessBlockInput) (*request.Request, *s3control.GetPublicAccessBlockOutput)

	GetStorageLensConfiguration(*s3control.GetStorageLensConfigurationInput) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationWithContext(aws.Context, *s3control.GetStorageLensConfigurationInput, ...request.Option) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationRequest(*s3control.GetStorageLensConfigurationInput) (*request.Request, *s3control.GetStorageLensConfigurationOutput)

	GetStorageLensConfigurationTagging(*s3control.GetStorageLensConfigurationTaggingInput) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	GetStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.GetStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	GetStorageLensConfigurationTaggingRequest(*s3control.GetStorageLensConfigurationTaggingInput) (*request.Request, *s3control.GetStorageLensConfigurationTaggingOutput)

	ListAccessPoints(*s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsWithContext(aws.Context, *s3control.ListAccessPointsInput, ...request.Option) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsRequest(*s3control.ListAccessPointsInput) (*request.Request, *s3control.ListAccessPointsOutput)

	ListAccessPointsPages(*s3control.ListAccessPointsInput, func(*s3control.ListAccessPointsOutput, bool) bool) error
	ListAccessPointsPagesWithContext(aws.Context, *s3control.ListAccessPointsInput, func(*s3control.ListAccessPointsOutput, bool) bool, ...request.Option) error

	ListAccessPointsForObjectLambda(*s3control.ListAccessPointsForObjectLambdaInput) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	ListAccessPointsForObjectLambdaWithContext(aws.Context, *s3control.ListAccessPointsForObjectLambdaInput, ...request.Option) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	ListAccessPointsForObjectLambdaRequest(*s3control.ListAccessPointsForObjectLambdaInput) (*request.Request, *s3control.ListAccessPointsForObjectLambdaOutput)

	ListAccessPointsForObjectLambdaPages(*s3control.ListAccessPointsForObjectLambdaInput, func(*s3control.ListAccessPointsForObjectLambdaOutput, bool) bool) error
	ListAccessPointsForObjectLambdaPagesWithContext(aws.Context, *s3control.ListAccessPointsForObjectLambdaInput, func(*s3control.ListAccessPointsForObjectLambdaOutput, bool) bool, ...request.Option) error

	ListJobs(*s3control.ListJobsInput) (*s3control.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *s3control.ListJobsInput, ...request.Option) (*s3control.ListJobsOutput, error)
	ListJobsRequest(*s3control.ListJobsInput) (*request.Request, *s3control.ListJobsOutput)

	ListJobsPages(*s3control.ListJobsInput, func(*s3control.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *s3control.ListJobsInput, func(*s3control.ListJobsOutput, bool) bool, ...request.Option) error

	ListMultiRegionAccessPoints(*s3control.ListMultiRegionAccessPointsInput) (*s3control.ListMultiRegionAccessPointsOutput, error)
	ListMultiRegionAccessPointsWithContext(aws.Context, *s3control.ListMultiRegionAccessPointsInput, ...request.Option) (*s3control.ListMultiRegionAccessPointsOutput, error)
	ListMultiRegionAccessPointsRequest(*s3control.ListMultiRegionAccessPointsInput) (*request.Request, *s3control.ListMultiRegionAccessPointsOutput)

	ListMultiRegionAccessPointsPages(*s3control.ListMultiRegionAccessPointsInput, func(*s3control.ListMultiRegionAccessPointsOutput, bool) bool) error
	ListMultiRegionAccessPointsPagesWithContext(aws.Context, *s3control.ListMultiRegionAccessPointsInput, func(*s3control.ListMultiRegionAccessPointsOutput, bool) bool, ...request.Option) error

	ListRegionalBuckets(*s3control.ListRegionalBucketsInput) (*s3control.ListRegionalBucketsOutput, error)
	ListRegionalBucketsWithContext(aws.Context, *s3control.ListRegionalBucketsInput, ...request.Option) (*s3control.ListRegionalBucketsOutput, error)
	ListRegionalBucketsRequest(*s3control.ListRegionalBucketsInput) (*request.Request, *s3control.ListRegionalBucketsOutput)

	ListRegionalBucketsPages(*s3control.ListRegionalBucketsInput, func(*s3control.ListRegionalBucketsOutput, bool) bool) error
	ListRegionalBucketsPagesWithContext(aws.Context, *s3control.ListRegionalBucketsInput, func(*s3control.ListRegionalBucketsOutput, bool) bool, ...request.Option) error

	ListStorageLensConfigurations(*s3control.ListStorageLensConfigurationsInput) (*s3control.ListStorageLensConfigurationsOutput, error)
	ListStorageLensConfigurationsWithContext(aws.Context, *s3control.ListStorageLensConfigurationsInput, ...request.Option) (*s3control.ListStorageLensConfigurationsOutput, error)
	ListStorageLensConfigurationsRequest(*s3control.ListStorageLensConfigurationsInput) (*request.Request, *s3control.ListStorageLensConfigurationsOutput)

	ListStorageLensConfigurationsPages(*s3control.ListStorageLensConfigurationsInput, func(*s3control.ListStorageLensConfigurationsOutput, bool) bool) error
	ListStorageLensConfigurationsPagesWithContext(aws.Context, *s3control.ListStorageLensConfigurationsInput, func(*s3control.ListStorageLensConfigurationsOutput, bool) bool, ...request.Option) error

	PutAccessPointConfigurationForObjectLambda(*s3control.PutAccessPointConfigurationForObjectLambdaInput) (*s3control.PutAccessPointConfigurationForObjectLambdaOutput, error)
	PutAccessPointConfigurationForObjectLambdaWithContext(aws.Context, *s3control.PutAccessPointConfigurationForObjectLambdaInput, ...request.Option) (*s3control.PutAccessPointConfigurationForObjectLambdaOutput, error)
	PutAccessPointConfigurationForObjectLambdaRequest(*s3control.PutAccessPointConfigurationForObjectLambdaInput) (*request.Request, *s3control.PutAccessPointConfigurationForObjectLambdaOutput)

	PutAccessPointPolicy(*s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyWithContext(aws.Context, *s3control.PutAccessPointPolicyInput, ...request.Option) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyRequest(*s3control.PutAccessPointPolicyInput) (*request.Request, *s3control.PutAccessPointPolicyOutput)

	PutAccessPointPolicyForObjectLambda(*s3control.PutAccessPointPolicyForObjectLambdaInput) (*s3control.PutAccessPointPolicyForObjectLambdaOutput, error)
	PutAccessPointPolicyForObjectLambdaWithContext(aws.Context, *s3control.PutAccessPointPolicyForObjectLambdaInput, ...request.Option) (*s3control.PutAccessPointPolicyForObjectLambdaOutput, error)
	PutAccessPointPolicyForObjectLambdaRequest(*s3control.PutAccessPointPolicyForObjectLambdaInput) (*request.Request, *s3control.PutAccessPointPolicyForObjectLambdaOutput)

	PutBucketLifecycleConfiguration(*s3control.PutBucketLifecycleConfigurationInput) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationWithContext(aws.Context, *s3control.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationRequest(*s3control.PutBucketLifecycleConfigurationInput) (*request.Request, *s3control.PutBucketLifecycleConfigurationOutput)

	PutBucketPolicy(*s3control.PutBucketPolicyInput) (*s3control.PutBucketPolicyOutput, error)
	PutBucketPolicyWithContext(aws.Context, *s3control.PutBucketPolicyInput, ...request.Option) (*s3control.PutBucketPolicyOutput, error)
	PutBucketPolicyRequest(*s3control.PutBucketPolicyInput) (*request.Request, *s3control.PutBucketPolicyOutput)

	PutBucketReplication(*s3control.PutBucketReplicationInput) (*s3control.PutBucketReplicationOutput, error)
	PutBucketReplicationWithContext(aws.Context, *s3control.PutBucketReplicationInput, ...request.Option) (*s3control.PutBucketReplicationOutput, error)
	PutBucketReplicationRequest(*s3control.PutBucketReplicationInput) (*request.Request, *s3control.PutBucketReplicationOutput)

	PutBucketTagging(*s3control.PutBucketTaggingInput) (*s3control.PutBucketTaggingOutput, error)
	PutBucketTaggingWithContext(aws.Context, *s3control.PutBucketTaggingInput, ...request.Option) (*s3control.PutBucketTaggingOutput, error)
	PutBucketTaggingRequest(*s3control.PutBucketTaggingInput) (*request.Request, *s3control.PutBucketTaggingOutput)

	PutBucketVersioning(*s3control.PutBucketVersioningInput) (*s3control.PutBucketVersioningOutput, error)
	PutBucketVersioningWithContext(aws.Context, *s3control.PutBucketVersioningInput, ...request.Option) (*s3control.PutBucketVersioningOutput, error)
	PutBucketVersioningRequest(*s3control.PutBucketVersioningInput) (*request.Request, *s3control.PutBucketVersioningOutput)

	PutJobTagging(*s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error)
	PutJobTaggingWithContext(aws.Context, *s3control.PutJobTaggingInput, ...request.Option) (*s3control.PutJobTaggingOutput, error)
	PutJobTaggingRequest(*s3control.PutJobTaggingInput) (*request.Request, *s3control.PutJobTaggingOutput)

	PutMultiRegionAccessPointPolicy(*s3control.PutMultiRegionAccessPointPolicyInput) (*s3control.PutMultiRegionAccessPointPolicyOutput, error)
	PutMultiRegionAccessPointPolicyWithContext(aws.Context, *s3control.PutMultiRegionAccessPointPolicyInput, ...request.Option) (*s3control.PutMultiRegionAccessPointPolicyOutput, error)
	PutMultiRegionAccessPointPolicyRequest(*s3control.PutMultiRegionAccessPointPolicyInput) (*request.Request, *s3control.PutMultiRegionAccessPointPolicyOutput)

	PutPublicAccessBlock(*s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockWithContext(aws.Context, *s3control.PutPublicAccessBlockInput, ...request.Option) (*s3control.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockRequest(*s3control.PutPublicAccessBlockInput) (*request.Request, *s3control.PutPublicAccessBlockOutput)

	PutStorageLensConfiguration(*s3control.PutStorageLensConfigurationInput) (*s3control.PutStorageLensConfigurationOutput, error)
	PutStorageLensConfigurationWithContext(aws.Context, *s3control.PutStorageLensConfigurationInput, ...request.Option) (*s3control.PutStorageLensConfigurationOutput, error)
	PutStorageLensConfigurationRequest(*s3control.PutStorageLensConfigurationInput) (*request.Request, *s3control.PutStorageLensConfigurationOutput)

	PutStorageLensConfigurationTagging(*s3control.PutStorageLensConfigurationTaggingInput) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	PutStorageLensConfigurationTaggingWithContext(aws.Context, *s3control.PutStorageLensConfigurationTaggingInput, ...request.Option) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	PutStorageLensConfigurationTaggingRequest(*s3control.PutStorageLensConfigurationTaggingInput) (*request.Request, *s3control.PutStorageLensConfigurationTaggingOutput)

	SubmitMultiRegionAccessPointRoutes(*s3control.SubmitMultiRegionAccessPointRoutesInput) (*s3control.SubmitMultiRegionAccessPointRoutesOutput, error)
	SubmitMultiRegionAccessPointRoutesWithContext(aws.Context, *s3control.SubmitMultiRegionAccessPointRoutesInput, ...request.Option) (*s3control.SubmitMultiRegionAccessPointRoutesOutput, error)
	SubmitMultiRegionAccessPointRoutesRequest(*s3control.SubmitMultiRegionAccessPointRoutesInput) (*request.Request, *s3control.SubmitMultiRegionAccessPointRoutesOutput)

	UpdateJobPriority(*s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobPriorityWithContext(aws.Context, *s3control.UpdateJobPriorityInput, ...request.Option) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobPriorityRequest(*s3control.UpdateJobPriorityInput) (*request.Request, *s3control.UpdateJobPriorityOutput)

	UpdateJobStatus(*s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error)
	UpdateJobStatusWithContext(aws.Context, *s3control.UpdateJobStatusInput, ...request.Option) (*s3control.UpdateJobStatusOutput, error)
	UpdateJobStatusRequest(*s3control.UpdateJobStatusInput) (*request.Request, *s3control.UpdateJobStatusOutput)
}

var _ S3ControlAPI = (*s3control.S3Control)(nil)
