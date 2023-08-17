package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("job", job, "started")
		time.Sleep(time.Second) // simulate expensive task
		results <- job * 4
		fmt.Println("job", job, "finished")
	}
}

func main() {
	const n_jobs = 5
	jobs := make(chan int, n_jobs)
	results := make(chan int, n_jobs)

	// start goroutines. Initially are blocked, because there are no jobs are in the channel
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < n_jobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs allowed

	/**
	* Collect all the results of the work. 
	* This also ensures that the worker goroutines have finished. 
	* An alternative way to wait for multiple goroutines is to use a WaitGroup.
	*/
	for j := 0; j < n_jobs; j++ {
		<-results
	}
}