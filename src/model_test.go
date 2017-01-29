package main

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
)

func TestUnitDecodeFrontendRequest(t *testing.T) {
	const jsonStream = `{"artifact-url": "s3://input/foo.tar.bz2", "output": "s3://output/", "instance-count": 2, "instance-type": "m4.large", "region": "us-east-1"}`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m FrontendRequest
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		if m.ArtifactURL != "s3://input/foo.tar.bz2" {
			t.Fatalf("unexpected value: %s", m.ArtifactURL)
		}
		if m.Output != "s3://output/" {
			t.Fatalf("unexpected value: %s", m.Output)
		}
		if m.InstanceCount != 2 {
			t.Fatalf("unexpected value: %d", m.InstanceCount)
		}
		if m.InstanceType != "m4.large" {
			t.Fatalf("unexpected value: %s", m.InstanceType)
		}
		if m.Region != "us-east-1" {
			t.Fatalf("unexpected value: %s", m.Region)
		}
	}
}

func TestUnitEncodeFrontendResponse(t *testing.T) {
	response := FrontendResponse{"123"}
	val, err := json.Marshal(response)

	if err != nil {
		t.Fatal(err)
	}

	str := string(val)
	if str != `{"id":"123"}` {
		t.Fatalf("unexpected value: %s", str)
	}
}
