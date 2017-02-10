package main

import (
	"fmt"
	"net/http"
	"time"
)

func job(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Job need a name", http.StatusBadRequest)
		return
	}

	delay, err := time.ParseDuration(r.URL.Query().Get("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay}
	JobQueue <- job
	fmt.Println("Job " + job.Name + " queued")
	w.WriteHeader(http.StatusCreated)
	return
}
