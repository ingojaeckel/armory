package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func startTestOnWorker(artifactURL string, simulationClass string, ipAddresses []*string, testID string) error {
	jsonBytes, err := json.Marshal(workerPutRequest{artifactURL, simulationClass, testID})
	if err != nil {
		return err
	}
	body := bytes.NewReader(jsonBytes)
	client := &http.Client{}

	for _, ipAddress := range ipAddresses {
		req, _ := http.NewRequest("PUT", fmt.Sprintf("http://%s:8080/rest/worker", *ipAddress), body)
		// Starts load test on the current worker instance
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			var errorResp ErrorResponse
			if err := parseJSON(resp.Body, &errorResp); err != nil {
				// Failed to parse error response. Just return the error code.
				return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
			}
			// Successfully parsed the error response. Include it in the returned error.
			return fmt.Errorf("Request failed. Status code: %d, Description: %s", resp.StatusCode, errorResp.Description)
		}
	}
	return nil
}
