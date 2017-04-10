package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func parseJSON(reader io.Reader, val interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &val)
}

func writeJSON(w http.ResponseWriter, status int, r interface{}) error {
	return WriteJSONAndLog(w, status, r, false)
}

// WriteJSONAndLog Send reply as JSON to client and optionally log it.
func WriteJSONAndLog(w http.ResponseWriter, status int, r interface{}, log bool) error {
	val, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if log {
		Log("http status: %d, response: %s", status, string(val))
	}
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, string(val))
	return nil
}
