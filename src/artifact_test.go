package main

import "testing"

func TestPrepareGatlingDirectory(t *testing.T) {
	if err := prepareGatlingDirectory("https://s3-us-west-1.amazonaws.com/my-loadtest-result-bucket/example.tar.bz2", "/Users/ingo/gatling"); err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrepareGatlingDirectory2(t *testing.T) {
	if err := prepareGatlingDirectory("https://s3-us-west-1.amazonaws.com/my-loadtest-result-bucket/example.tar.bz2", "/Users/ingo/gatling"); err != nil {
		t.Errorf(err.Error())
	}

	if err := executeGatling("com.example.Loadtest"); err != nil {
		t.Errorf(err.Error())
	}
}
