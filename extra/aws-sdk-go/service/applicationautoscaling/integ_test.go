// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.18 && integration
// +build go1.18,integration

package applicationautoscaling_test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting/integration"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeScalableTargets(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := applicationautoscaling.New(sess)
	params := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: aws.String("ec2"),
	}
	_, err := svc.DescribeScalableTargetsWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
