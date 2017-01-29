package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

type Test struct {
	ID string
}

type TestManager interface {
	startTest(r FrontendRequest) (bool, ErrorResponse)
}

func NewTest() Test {
	return Test{uuid.NewV4().String()}
}

func (t Test) startTest(r FrontendRequest) (bool, ErrorResponse) {
	if conf.Base.SkipInstanceCreation {
		updateStateWithDebug(t.ID, "Skipping instance creation", "2_skip_instance_creation")
		updateStateWithDebug(t.ID, "Starting load test on worker instances", "3_start_test_on_worker_instances")

		fallbackAddresses := []*string{&conf.Base.FallbackWorkerIPAddress}

		if err := startTestOnWorker(r.ArtifactURL, r.SimulationClass, fallbackAddresses, t.ID); err != nil {
			Log(err.Error())
			return false, ErrorResponse{fmt.Sprintf("Failed to start test on worker instance(s): %s", err.Error())}
		}
		updateStateWithDebug(t.ID, "Started test on worker(s)", "4_test_running")
		return true, ErrorResponse{}
	}

	runner := AwsInstanceRunner{GetAwsService(r.Region)}

	// Create worker instances
	updateStateWithDebug(t.ID, "Creating worker instances", "1_creating_worker_instances")
	createResp, err := runner.CreateInstances(StartInstancesRequest{
		InstanceType:    r.InstanceType,
		Region:          r.Region,
		Count:           r.InstanceCount,
		AmiID:           conf.Base.AmiID,
		SubnetID:        conf.Base.VpcSubnet,
		SecurityGroupID: conf.Base.SecurityGroup,
	})
	if err != nil {
		Log(err.Error())
		return false, ErrorResponse{"Unable to create instance(s)"}
	}

	updateStateWithDebug(t.ID, "Waiting until instances are running", "2_wait_until_instances_are_running")
	runningResp, err := runner.WaitUntilRunning(r.Region, createResp.InstanceIds)
	if err != nil {
		Log(err.Error())
		return false, ErrorResponse{"Failed waiting for instance(s) to start running"}
	}

	updateStateWithDebug(t.ID, "Starting load test on worker instances", "3_start_test_on_worker_instances")
	if err := startTestOnWorker(r.ArtifactURL, r.SimulationClass, runningResp.IPAddresses, t.ID); err != nil {
		Log(err.Error())
		return false, ErrorResponse{"Failed to start test on worker instance(s)"}
	}
	updateStateWithDebug(t.ID, "Started test on worker(s)", "4_test_running")
	runner.TerminateInstances(r.Region, createResp.InstanceIds)

	return true, ErrorResponse{}
}
