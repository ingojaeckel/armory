package main

import (
	"fmt"
	"net/http"
)

func handleTestPut(w http.ResponseWriter, r *http.Request) {
	var putRequest FrontendRequest
	if err := parseJSON(r.Body, &putRequest); err != nil {
		WriteJSONAndLog(w, 400, ErrorResponse{fmt.Sprintf("Invalid request: %s", err.Error())}, true)
		return
	}

	test := NewTest()
	updateStateWithDebug(test.ID, "Parsed input", "parsed_input")
	// Start the test in the background
	go func() {
		if success, errResponse := test.startTest(putRequest); !success {
			updateStateWithDebug(test.ID, fmt.Sprintf("Failed to start the test: %s", errResponse.Description), "failed")
			return
		}
		updateStateWithDebug(test.ID, fmt.Sprintf("Started test successfully"), "started_test")
	}()

	// send back an OK with the test ID. clients can retrieve the test status via the GET /rest/test?id=.. API.
	writeJSON(w, 201, FrontendResponse{test.ID})
}

func handleTestGet(w http.ResponseWriter, r *http.Request) {
	testID := r.URL.Query().Get("id")
	state := getState(testID)
	writeJSON(w, 200, testStatusResponse{state})
}

func handleWorkerPut(w http.ResponseWriter, r *http.Request) {
	var putRequest workerPutRequest
	if err := parseJSON(r.Body, &putRequest); err != nil {
		WriteJSONAndLog(w, 400, ErrorResponse{fmt.Sprintf("Invalid request: %s", err.Error())}, true)
		return
	}

	updateStateWithDebug(putRequest.TestID, "Preparing gatling directory", "preparing_gatling_directoy")
	if err := prepareGatlingDirectory(putRequest.ArtifactURL, conf.Base.GatlingRoot); err != nil {
		WriteJSONAndLog(w, 500, ErrorResponse{fmt.Sprintf("Failed to download artifact from %s: %s", putRequest.ArtifactURL, err.Error())}, true)
		return
	}

	updateStateWithDebug(putRequest.TestID, "Starting load test..", "starting_test")
	if err := executeGatling(putRequest.SimulationClass); err != nil {
		WriteJSONAndLog(w, 500, ErrorResponse{fmt.Sprintf("Failed to execute gatling: %s", err.Error())}, true)
		return
	}

	// TODO upload results to s3
	// TODO get the s3 bucket & key from workerPutRequest
	// uploadToS3(s3bucket string, s3key string)
	updateStateWithDebug(putRequest.TestID, "Done running test. Will terminate.", "done")

	if err := execute([]string{}, "/usr/bin/sudo", "/sbin/shutdown", "-h", "now"); err != nil {
		WriteJSONAndLog(w, 500, ErrorResponse{fmt.Sprintf("Failed to terminate worker instance: %s", err.Error())}, true)
		return
	}
	updateStateWithDebug(putRequest.TestID, "Terminated", "terminated.")
}
