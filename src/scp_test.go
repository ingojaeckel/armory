package main

import "testing"

func TestScp(t *testing.T) {
	ip := "54.86.70.47"
	if err := executeScp("ec2-user", ip, "../infra/packer/worker/terraform.pem", "scp.go", "/tmp"); err != nil {
		t.Errorf("Failed to run scp: %s", err.Error())
	}
}
