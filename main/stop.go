package main

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Stop struct {
	fs         *flag.FlagSet
	response   string
	instanceID string
	sess       *session.Session
}

func (stop *Stop) Name() string {
	return stop.fs.Name()
}

func (stop *Stop) Init(args []string) error {
	return stop.fs.Parse(args)
}

func (stop *Stop) Run() string {
	stop.sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(stop.sess)
	// Dry run
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(stop.instanceID),
		},
		DryRun: aws.Bool(true),
	}
	_, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if !ok || awsErr.Code() != "DryRunOperation" {
		stop.response = "AWS error: " + awsErr.Code()
		return stop.response
	}
	input.DryRun = aws.Bool(false)
	_, err = svc.StopInstances(input)
	if err != nil {
		stop.response = err.Error()
		return stop.response
	}
	stop.response = "Instance " + stop.instanceID + " stopped."
	return stop.response
}
