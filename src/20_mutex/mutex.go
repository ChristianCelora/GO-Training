package main

import (
	"fmt"
	"sync"
)

func doIncrement(c *Container, wg *sync.WaitGroup, key string, n int) {
	for i := 0; i < n; i++ {
		c.inc(key)
	}
	wg.Done()
}

type Container struct {
	mu sync.Mutex	// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
	counters map[string]int
}

func (c *Container) inc(key string){
	// Lock the mutex before accessing counters
	c.mu.Lock()
	// unlock it at the end of the function using a defer statement.
	defer c.mu.Unlock()
	c.counters[key]++
}

func main(){
	/**
	* In the previous example we saw how to manage simple counter state using atomic operations. 
	* For more complex state we can use a mutex to safely access data across multiple goroutines.
	*/
	c := Container {
		// Note: a zero value mutex is usable as-is. No initialization is required
		counters: map[string]int {"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

		
    /*doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }*/

	wg.Add(3)
	go doIncrement(&c, &wg, "a", 1000)
	go doIncrement(&c, &wg, "a", 1000)
	go doIncrement(&c, &wg, "b", 1000)
	/*go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)*/
	wg.Wait()
	fmt.Println(c.counters)
}