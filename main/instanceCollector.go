package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Instance struct {
	InstanceID   string
	Name         string
	InstanceType string
	State        string
	Region       string
}

type collector struct {
	allInstances []ec2Instance
	sess         *session.Session
}

func newCollector() *collector {
	return &collector{
		sess: session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))}
}

func (c *collector) collectAll() []ec2Instance {
	var allInstances []ec2Instance
	for _, region := range allRegions {
		instancesInRegion := c.collectInRegion(region)
		allInstances = append(allInstances, instancesInRegion...)
	}
	return allInstances
}

func (c *collector) collectInRegion(region string) []ec2Instance {
	ec2Client := ec2.New(c.sess, &aws.Config{ // client for ec2
		Region: aws.String(region),
	})
	// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#Filter
	// filter := []*ec2.Filter{
	// 	&ec2.Filter{
	// 		Name:   aws.String("instance-state-name"),
	// 		Values: []*string{aws.String("pending"), aws.String("running")},
	// 	},
	// }
	filter := []*ec2.Filter{
		&ec2.Filter{},
	}
	// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput
	input := ec2.DescribeInstancesInput{
		Filters: filter,
	}
	describeInstancesOutput, _ := ec2Client.DescribeInstances(&input)

	// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesOutput
	var instancesInRegion []ec2Instance
	for _, res := range describeInstancesOutput.Reservations {
		// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#Instance
		for _, inst := range res.Instances {
			name := getSpecificTagValue("Name", inst.Tags)
			instanceAttribute := ec2Instance{*inst.InstanceId, name, *inst.InstanceType, *inst.State.Name, region}
			instancesInRegion = append(instancesInRegion, instanceAttribute)
		}
	}
	return instancesInRegion
}

func getSpecificTagValue(key string, tags []*ec2.Tag) string {
	for _, tag := range tags {
		if *(tag.Key) == key {
			return *tag.Value
		}
	}
	return "--"
}
