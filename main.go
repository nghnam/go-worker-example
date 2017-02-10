package main

import (
	"net/http"
)

// JobQueue ...
var JobQueue chan Job

func main() {
	JobQueue = make(chan Job, 5)

	dispatcher := NewDispatcher(2)
	dispatcher.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/job", job)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
