package main

import (
	"fmt"
	"time"
)

// Worker ...
type Worker struct {
	ID         int
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker ...
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		ID:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start ...
func (w Worker) Start() {
	go func() {
		for {
			// register current to worker workerPool
			// this is channel to communicate from dispatcher to worker
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel: // receive job
				fmt.Printf("Worker %d: received job %s, delay for %f seconds\n", w.ID, job.Name, job.Delay.Seconds())
				time.Sleep(job.Delay)
				fmt.Printf("Worker %d: job %s done\n", w.ID, job.Name)

			case <-w.quit: // receive quit signal
				return
			}
		}
	}()
}

// Stop ...
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
