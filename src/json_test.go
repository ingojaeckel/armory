package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestUnitParseJson(t *testing.T) {
	var r workerPutRequest

	if err := parseJSON(getReadCloser(`{"artifact-url": "some-url", "simulation-class": "Foo"}`), &r); err != nil {
		t.Error(err.Error())
	}
	if r.ArtifactURL != "some-url" {
		t.Errorf("Unexpected value: %s", r.ArtifactURL)
	}
	if r.SimulationClass != "Foo" {
		t.Errorf("Unexpected value: %s", r.SimulationClass)
	}
}

func TestUnitParseMalformedJson(t *testing.T) {
	var r workerPutRequest

	if err := parseJSON(getReadCloser(`{"artifact-url":`), &r); err == nil {
		t.Error("This should have failed due to malformed JSON")
	}
}

func getReadCloser(str string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(str)))
}
