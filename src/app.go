package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, "up")
	})
	http.HandleFunc("/rest/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			handleTestPut(w, r)
		} else if r.Method == "GET" {
			handleTestGet(w, r)
		} else {
			w.WriteHeader(415)
		}
	})
	http.HandleFunc("/rest/worker", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			handleWorkerPut(w, r)
		} else {
			w.WriteHeader(415)
		}
	})
	http.HandleFunc("/rest/version", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			io.WriteString(w, Version)
		} else {
			w.WriteHeader(415)
		}
	})
	if err := initConfiguration(); err != nil {
		fmt.Printf("Failed to intialize the configuration: %s", err.Error())
		return
	}
	Log("Running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
