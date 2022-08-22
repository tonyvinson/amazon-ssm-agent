// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2instanceconnect_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To push an SSH key to an EC2 instance
//
// The following example pushes a sample SSH public key to the EC2 instance i-abcd1234
// in AZ us-west-2b for use by the instance OS user ec2-user.
func ExampleEC2InstanceConnect_SendSSHPublicKey_shared00() {
	svc := ec2instanceconnect.New(session.New())
	input := &ec2instanceconnect.SendSSHPublicKeyInput{
		AvailabilityZone: aws.String("us-west-2a"),
		InstanceId:       aws.String("i-abcd1234"),
		InstanceOSUser:   aws.String("ec2-user"),
		SSHPublicKey:     aws.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC3FlHqj2eqCdrGHuA6dRjfZXQ4HX5lXEIRHaNbxEwE5Te7xNF7StwhrDtiV7IdT5fDqbRyGw/szPj3xGkNTVoElCZ2dDFb2qYZ1WLIpZwj/UhO9l2mgfjR56UojjQut5Jvn2KZ1OcyrNO0J83kCaJCV7JoVbXY79FBMUccYNY45zmv9+1FMCfY6i2jdIhwR6+yLk8oubL8lIPyq7X+6b9S0yKCkB7Peml1DvghlybpAIUrC9vofHt6XP4V1i0bImw1IlljQS+DUmULRFSccATDscCX9ajnj7Crhm0HAZC0tBPXpFdHkPwL3yzYo546SCS9LKEwz62ymxxbL9k7h09t"),
	}

	result, err := svc.SendSSHPublicKey(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ec2instanceconnect.ErrCodeAuthException:
				fmt.Println(ec2instanceconnect.ErrCodeAuthException, aerr.Error())
			case ec2instanceconnect.ErrCodeInvalidArgsException:
				fmt.Println(ec2instanceconnect.ErrCodeInvalidArgsException, aerr.Error())
			case ec2instanceconnect.ErrCodeServiceException:
				fmt.Println(ec2instanceconnect.ErrCodeServiceException, aerr.Error())
			case ec2instanceconnect.ErrCodeThrottlingException:
				fmt.Println(ec2instanceconnect.ErrCodeThrottlingException, aerr.Error())
			case ec2instanceconnect.ErrCodeEC2InstanceNotFoundException:
				fmt.Println(ec2instanceconnect.ErrCodeEC2InstanceNotFoundException, aerr.Error())
			case ec2instanceconnect.ErrCodeEC2InstanceStateInvalidException:
				fmt.Println(ec2instanceconnect.ErrCodeEC2InstanceStateInvalidException, aerr.Error())
			case ec2instanceconnect.ErrCodeEC2InstanceUnavailableException:
				fmt.Println(ec2instanceconnect.ErrCodeEC2InstanceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
