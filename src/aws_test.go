package main

import "testing"

func TestAwsCreateInstance(t *testing.T) {
	region := "us-east-1"
	runner := AwsInstanceRunner{GetAwsService(region)}
	resp, err := runner.CreateInstances(StartInstancesRequest{"m4.large", region, "sg-faac829e", "ami-31490d51", "subnet-9f25fbfb", 2})

	if err != nil {
		t.Errorf("Unable to create instance: %s", err.Error())
	}

	if err := runner.TerminateInstances(region, resp.InstanceIds); err != nil {
		t.Errorf("Unable to terminate instance: %s", err.Error())
	}
}

func TestAwsListInstances(t *testing.T) {
	region := "us-west-1"
	runner := AwsInstanceRunner{GetAwsService(region)}
	if err := runner.ListInstances(region); err != nil {
		t.Errorf("Unable to create instance: %s", err.Error())
	}
}
