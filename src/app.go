package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	if err := initConfiguration(); err != nil {
		fmt.Printf("Failed to intialize the configuration: %s", err.Error())
		return
	}
	confStr, _ := toJSON(conf)
	fmt.Printf("Using config: %s\n", string(confStr))

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/health"), func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, Version)
	})
	mux.HandleFunc(pat.Get("/rest/version"), func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, Version)
	})
	mux.HandleFunc(pat.Get("/rest/test"), handleTestGet)
	mux.HandleFunc(pat.Put("/rest/test"), handleTestPut)
	mux.HandleFunc(pat.Put("/rest/worker"), handleWorkerPut)

	Log("Running..")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
