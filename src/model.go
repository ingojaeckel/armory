package main

// FrontendRequest A request for the front-end.
type FrontendRequest struct {
	ArtifactURL     string `json:"artifact-url"`
	SimulationClass string `json:"simulation-class"`
	Output          string `json:"output"`
	Region          string `json:"region"`
	InstanceCount   int64  `json:"instance-count"`
	InstanceType    string `json:"instance-type"`
}

// FrontendResponse The response object sent back by the front-end.
type FrontendResponse struct {
	ID string `json:"id"`
}

type testStatusResponse struct {
	Status string `json:"status"`
}

type workerPutRequest struct {
	ArtifactURL     string `json:"artifact-url"`
	SimulationClass string `json:"simulation-class"`
	TestID          string `json:"test-id"`
}

type workerGetResponse struct {
	Status string `json:"status"`
}

// ErrorResponse The generic error response that will be send to clients.
type ErrorResponse struct {
	Description string `json:"description"`
}
