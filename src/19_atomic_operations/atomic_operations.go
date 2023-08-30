package main 

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/**
	* There are a few other options for managing state though. 
	* Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
	*/
	var atomic_counter uint64
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1);
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&atomic_counter, 1)
				counter++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("non-atomic counter", counter)
	fmt.Println("atomic counter", atomic_counter)
}