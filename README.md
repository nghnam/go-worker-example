# Example Worker in Golang

# Source:
    http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
    http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html

# Test
```
$ go build
$ ./go-worker-example
$ for i in {1..10}; do curl "http://localhost:8080/job?name=job$i&delay=3s"; done
```

# Log
```
Job job1 queued
Worker 1: received job job1, delay for 3.000000 seconds
Job job2 queued
Worker 0: received job job2, delay for 3.000000 seconds
Job job3 queued
Job job4 queued
Job job5 queued
Job job6 queued
Job job7 queued
Job job8 queued
Job job9 queued
Job job10 queued
Worker 1: job job1 done
Worker 1: received job job3, delay for 3.000000 seconds
Worker 0: job job2 done
Worker 0: received job job4, delay for 3.000000 seconds
Worker 1: job job3 done
Worker 1: received job job5, delay for 3.000000 seconds
Worker 0: job job4 done
Worker 0: received job job6, delay for 3.000000 seconds
Worker 1: job job5 done
Worker 1: received job job7, delay for 3.000000 seconds
Worker 0: job job6 done
Worker 0: received job job8, delay for 3.000000 seconds
Worker 1: job job7 done
Worker 1: received job job9, delay for 3.000000 seconds
Worker 0: job job8 done
Worker 0: received job job10, delay for 3.000000 seconds
Worker 1: job job9 done
Worker 0: job job10 done
```
