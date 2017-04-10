package main

import (
	"log"
	"net/http"
)

func main() {
	Log("Running..")
	log.Fatal(http.ListenAndServe(":8080", setupMux()))
}
