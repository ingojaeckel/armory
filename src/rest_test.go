package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUnitTestGet(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test?id=123", nil)
	handleTestGet(rr, req)
	if rr.Body.String() != `{"status":"unknown"}` {
		t.Errorf("Unexpected status: %s", rr.Body.String())
	}
}

func TestAwsIntegrationTest(t *testing.T) {
	local := false

	var ip string
	var port int32

	if local {
		fmt.Println("Using local configuration")
		ip = "127.0.0.1"
		port = 8080
	} else {
		fmt.Println("Using remote configuration")
		ip = "54.86.54.1"
		port = 8080
	}

	url := fmt.Sprintf("http://%s:%d/rest/test", ip, port)
	fr := FrontendRequest{
		ArtifactURL:     "https://s3-us-west-1.amazonaws.com/my-loadtest-result-bucket/example.tar.bz2",
		Output:          "",
		InstanceCount:   1,
		InstanceType:    "t2.micro",
		Region:          "us-east-1",
		SimulationClass: "com.example.Loadtest",
	}
	b, err := toJSON(fr)
	if err != nil {
		t.Error(err.Error())
	}

	r, _ := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	r.Header.Set("Content-Type", "application/json")
	c := &http.Client{}
	httpResponse, err := c.Do(r)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("status: %s\n", httpResponse.Status)
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode >= 400 {
		var errorResp ErrorResponse
		parseJSON(httpResponse.Body, &errorResp)
		t.Errorf("Request failed with description: %s; status: %d", errorResp.Description, httpResponse.StatusCode)
	}

	var frontendResponse FrontendResponse
	if err := parseJSON(httpResponse.Body, &frontendResponse); err != nil {
		t.Error(err.Error())
	}
	if frontendResponse.ID == "" {
		t.Error("test id must not be empty")
	}
	fmt.Printf("Started test: %s\n", frontendResponse.ID)

	// hit the GET endpoint until the test is confirmed to be running
	for i := 0; i < 60; i++ {
		getResp, _ := http.Get(fmt.Sprintf("http://%s:%d/rest/test?id=%s", ip, port, frontendResponse.ID))

		var statResp testStatusResponse
		if err := parseJSON(getResp.Body, &statResp); err != nil {
			t.Error(err.Error())
		}

		fmt.Printf("http-status: %d; description: %s\n", getResp.StatusCode, statResp.Status)
		time.Sleep(time.Duration(10) * time.Second)
	}
}
