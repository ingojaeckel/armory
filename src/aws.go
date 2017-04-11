package main

import (
	"fmt"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

type StartInstancesRequest struct {
	InstanceType, Region, SecurityGroupID, AmiID, SubnetID string
	Count                                                  int64
}

type StartInstancesResponse struct {
	InstanceIds []*string
}

type RunningInstancesResponse struct {
	InstanceIds []*string
	IPAddresses []*string
}

type InstanceRunner interface {
	CreateInstances(r StartInstancesRequest) (StartInstancesResponse, error)
	WaitUntilRunning(region string, instanceIds []*string) (RunningInstancesResponse, error)
	TerminateInstances(region string, instanceIds []*string) error
	ListInstances(region string) error
}

type AwsInstanceRunner struct {
	svc *ec2.EC2
}

func (runner AwsInstanceRunner) CreateInstances(r StartInstancesRequest) (StartInstancesResponse, error) {
	in := &ec2.RunInstancesInput{
		ImageId:            aws.String(r.AmiID),
		MinCount:           aws.Int64(r.Count),
		MaxCount:           aws.Int64(r.Count),
		InstanceType:       aws.String(r.InstanceType),
		SecurityGroupIds:   []*string{aws.String(r.SecurityGroupID)},
		SubnetId:           aws.String(r.SubnetID),
		KeyName:            aws.String(conf.Base.Keyname),
		IamInstanceProfile: conf.IamInstanceProfile,
	}

	runInstancesResp, err := runner.svc.RunInstances(in)
	if err != nil {
		return StartInstancesResponse{}, err
	}

	// resp has all of the response data, pull out instance IDs:
	instanceIds := getInstanceIds(runInstancesResp)
	createTags := &ec2.CreateTagsInput{
		Resources: instanceIds,
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("Worker"),
			},
		},
		DryRun: aws.Bool(false),
	}

	_, err = runner.svc.CreateTags(createTags)
	return StartInstancesResponse{getInstanceIds(runInstancesResp)}, err
}

func (runner AwsInstanceRunner) WaitUntilRunning(region string, instanceIds []*string) (RunningInstancesResponse, error) {
	params := &ec2.DescribeInstancesInput{InstanceIds: instanceIds}

	err := retry(true, 1000, 15*1000, 30, func() (bool, error) {
		Log("Waiting for all instances to be running.")
		resp, err := runner.svc.DescribeInstances(params)

		if err != nil {
			return false, err
		}

		runningInstancesCount := 0
		Log("Checking response: %s", *resp)
		for idx := range resp.Reservations {
			for _, inst := range resp.Reservations[idx].Instances {
				Log("Instance ID: %s; State: %s", *inst.InstanceId, *inst.State.Name)
				if ec2.InstanceStateNameRunning == *inst.State.Name {
					runningInstancesCount++
				}
			}
		}
		// Are all instances running?
		return runningInstancesCount == len(instanceIds), nil
	})
	if err != nil {
		return RunningInstancesResponse{}, err
	}

	// Describe instances one more time to determine their public IP which should be set by now.
	describeResponse, _ := runner.svc.DescribeInstances(params)
	runningResp := RunningInstancesResponse{getInstanceIds(describeResponse.Reservations[0]), getPublicIPAddresses(describeResponse.Reservations[0])}

	for _, ipAddress := range runningResp.IPAddresses {
		if err := anyError(
			WaitForOpenPort(*ipAddress, 22, 15),
			// Copy latest version of executable to worker. Attempt up to 3 times.
			scpWithRetries(conf.Base.WorkerUsername, *ipAddress, "/tmp/terraform.pem", "/tmp/app", "/tmp/app"),
			// Start app on worker instance
			executeSSH(conf.Base.WorkerUsername, *ipAddress, "/tmp/terraform.pem", "/bin/bash /tmp/start.sh"),
			WaitForOpenPort(*ipAddress, 8080, 5)); err != nil {
			return RunningInstancesResponse{}, err
		}
	}
	return runningResp, nil
}

func anyError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (runner AwsInstanceRunner) TerminateInstances(region string, instanceIds []*string) error {
	resp, err := runner.svc.TerminateInstances(&ec2.TerminateInstancesInput{InstanceIds: instanceIds})
	Log("> TerminateInstances response: %s", resp)
	return err
}

func (runner AwsInstanceRunner) ListInstances(region string) error {
	resp, err := runner.svc.DescribeInstances(nil)
	if err != nil {
		return err
	}

	// resp has all of the response data, pull out instance IDs:
	Log("> Number of reservation sets: %d\n", len(resp.Reservations))
	for idx, res := range resp.Reservations {
		Log("  > Number of instances: %d\n", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			Log("    - Instance ID: %s, State: %s\n", *inst.InstanceId, *inst.State.Name)
		}
	}
	return nil
}

func getInstanceIds(resp *ec2.Reservation) []*string {
	ids := make([]*string, len(resp.Instances))
	for i := 0; i < len(resp.Instances); i++ {
		ids[i] = resp.Instances[i].InstanceId
	}
	return ids
}

func WaitForOpenPort(ipAddress string, port int32, maxAttempts int) error {
	address := fmt.Sprintf("%s:%d", ipAddress, port)
	Log("Waiting for open port: %s", address)

	return retry(true, 1000, 16*1000, maxAttempts, func() (bool, error) {
		if _, err := net.Dial("tcp", address); err != nil {
			return false, err
		}
		// Was able to connect to this port without receiving an error. Looks like this port is open.
		return true, nil
	})
}

func getPublicIPAddresses(resp *ec2.Reservation) []*string {
	ids := make([]*string, len(resp.Instances))
	for i := 0; i < len(resp.Instances); i++ {
		ids[i] = resp.Instances[i].PublicIpAddress
	}
	return ids
}

func GetAwsService(region string) *ec2.EC2 {
	return ec2.New(session.New(), getAwsConfig(region))
}

func GeAwsS3Service(region string) *s3.S3 {
	return s3.New(session.New(), getAwsConfig(region))
}

func getAwsConfig(region string) *aws.Config {
	return &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(conf.Base.AwsAccessKey, conf.Base.AwsSecretKey, ""),
	}
}
