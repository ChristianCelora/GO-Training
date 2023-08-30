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
	// Note: The order of workers starting up and finishing is likely to be different for each invocation.

	/**
	* Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service
	*/
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter // limit a request each 200ms
		fmt.Println("request", req, time.Now())
	}

	/**
	* We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. 
	* We can accomplish this by buffering our limiter channel. 
	* This burstLimiter channel will allow bursts of up to 3 events
	*/
	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		// the first 3 signals from the channel are received instantly. The other will be set with a ticker of 200ms
		// the value time.Now is irrelevant (can be any time.Time value)
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	// Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstLimiter.
	for req := range burstRequests {
        <-burstLimiter
        fmt.Println("request", req, time.Now())
    }
}