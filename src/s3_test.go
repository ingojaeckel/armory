package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestS3Upload(t *testing.T) {
	s3bucket := "my-loadtest-result-bucket"
	s3key := "hello_world.txt"
	payload, _ := ioutil.ReadFile("s3.go")
	reader := bytes.NewReader(payload)

	if err := uploadToS3(reader, s3bucket, s3key); err != nil {
		t.Errorf("failed to upload to s3: %s", err.Error())
	}
}
