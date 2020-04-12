package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Start struct {
	fs         *flag.FlagSet
	response   string
	instanceID string
	sess       *session.Session
}

func (start *Start) Name() string {
	return start.fs.Name()
}

func (start *Start) Init(args []string) error {
	return start.fs.Parse(args)
}

func (start *Start) Run() string {
	start.sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(start.sess)
	// Dry run
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(start.instanceID),
		},
		DryRun: aws.Bool(true),
	}
	_, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)
	if !ok && awsErr.Code() != "DryRunOperation" {
		start.response = "AWS error: " + awsErr.Code()
		return start.response
	}
	input.DryRun = aws.Bool(false)
	_, err = svc.StartInstances(input)
	if err != nil {
		start.response = err.Error()
		return start.response
	}
	start.response = "Instance " + start.instanceID + " started."
	return start.response
}
