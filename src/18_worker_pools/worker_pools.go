package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("job", job, "started")
		time.Sleep(time.Second) // simulate expensive task
		results <- job * 4
		fmt.Println("job", job, "finished")
	}
}

func simple_worker(id int) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Second) // simulate expensive task
	fmt.Println("worker", id, "finished")
}

func main() {
	const n_worker = 3
	const n_jobs = 5
	jobs := make(chan int, n_jobs)
	results := make(chan int, n_jobs)

	// start goroutines. Initially are blocked, because there are no jobs are in the channel
	for i := 0; i < n_worker; i++ {
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

	/**
	* Wait Groups
	*/
	var wg sync.WaitGroup

	for i := 0; i < n_worker; i++ {
		wg.Add(1)

		// Avoid use i inside each goroutine closure. Also passing i to the function as parmeter is fine.
		idx := i

		go func () {
			defer wg.Done()
			// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
			simple_worker(idx)
		} ()

		// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
		wg.Wait()
	}
}