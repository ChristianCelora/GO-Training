package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	* Select
	*
	* Go’s select lets you wait on multiple channel operations. 
	* Combining goroutines and channels with select is a powerful feature of Go.
	*/

	c1 := make(chan string)
	c2 := make(chan string)

	// Note: the total execution time of these 2 goroutines is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
	go func(c1 chan string) {
		time.Sleep(time.Second * 1)
		c1 <- "1 second has passed"
	} (c1)

	go func(c2 chan string) {
		time.Sleep(time.Second * 2)
		c2 <- "2 seconds has passed"
	} (c2)

	/**
	* Timeouts
	*
	* Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
	*/
	c3 := make(chan string, 1)
	go func(c3 chan string) {
		time.Sleep(time.Second * 4)
		c3 <- "4 seconds has passed"
	} (c3)

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 3; i++ {
		select {
		case msg := <-c1:
			fmt.Println("received", msg)
		case msg := <-c2:
			fmt.Println("received", msg)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout after 3 seconds")
		}		
	}
}