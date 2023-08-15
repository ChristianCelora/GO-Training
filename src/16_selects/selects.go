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

	/**
	* Non-blocking Channel Operations 
	*
	* Basic sends and receives on channels are blocking. 
	* However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
	*/
	messages := make(chan string)
	signals := make(chan string)

	/**
	* Here’s a non-blocking receive. 
	* If a value is available on messages then select will take the <-messages case with that value. 
	* If not it will immediately take the default case.
	*/
	// Uncomment below to send a message to the channel and select the first case
	go func(messages chan string) { messages <- "hello" } (messages)
	time.Sleep(time.Second)

	select {
	case msg := <-messages:
		fmt.Println("message received:", msg)
	default:
		fmt.Println("no message received")
	}

	/**
	* A non-blocking send works similarly. 
	* Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. 
	* Therefore the default case is selected.
	*
	* If messages has a buffer, messages := make(chan string, 1) the first case is selected.
	*/
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent:", msg)
	default:
		fmt.Println("no message sent")
	}

	/**
	* We can use multiple cases above the default clause to implement a multi-way non-blocking select. 
	* Here we attempt non-blocking receives on both messages and signals.
	*/
	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	case sig := <-signals:
		fmt.Println("signals received", sig)
	default:
		fmt.Println("no activity")
	}
}