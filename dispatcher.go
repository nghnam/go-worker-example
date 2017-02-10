package main

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorker  int
}

// NewDispatcher ...
func NewDispatcher(maxWorker int) *Dispatcher {
	pool := make(chan chan Job, maxWorker)
	return &Dispatcher{
		WorkerPool: pool,
		MaxWorker:  maxWorker,
	}
}

// Run ...
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// get idle worker
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// send job to worker
				jobChannel <- job
			}(job)
		}
	}
}
