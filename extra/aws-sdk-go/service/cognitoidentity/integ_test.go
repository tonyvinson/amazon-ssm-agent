// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.18 && integration
// +build go1.18,integration

package cognitoidentity_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/integration"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListIdentityPools(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cognitoidentity.New(sess)
	params := &cognitoidentity.ListIdentityPoolsInput{
		MaxResults: aws.Int64(10),
	}
	_, err := svc.ListIdentityPoolsWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_DescribeIdentityPool(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := cognitoidentity.New(sess)
	params := &cognitoidentity.DescribeIdentityPoolInput{
		IdentityPoolId: aws.String("us-east-1:aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	}
	_, err := svc.DescribeIdentityPoolWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(aerr.Message()) == 0 {
		t.Errorf("expect non-empty error message")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
