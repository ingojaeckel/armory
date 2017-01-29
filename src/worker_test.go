package main

import (
	"testing"
)

func TestStartWorker(t *testing.T) {
	host := "localhost"
	startTestOnWorker("https://s3-us-west-1.amazonaws.com/my-loadtest-result-bucket/example.tar.bz2", "com.example.Loadtest", []*string{&host}, "1")
}
