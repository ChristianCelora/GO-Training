package main

import (
	"fmt"
	"time"
)

func delayEcho(msg string) {
	time.Sleep(time.Second)
	fmt.Println(msg)
}

func f(from string, index int) {
    fmt.Println(from, ":", index)
}

func getChallenge(channel chan string) { 
	channel <- "test" 
}

func main() {
	/**
	* A goroutine is a lightweight thread of execution.
	* 
	* To invoke this function in a goroutine, use go f(s). 
	* This new goroutine will execute concurrently with the calling one.
	*/
	go delayEcho("delayed hello world")

    for i := 0; i < 3; i++ {
		go f("goroutine", i)
	}

	// You can also start a goroutine for an anonymous function call.
	go func (msg string) {
		fmt.Println(msg)
	} ("hello world")

	/**
	* Channels are the pipes that connect concurrent goroutines. 
	* You can send values into channels from one goroutine and receive those values into another goroutine.
	*/
	channel := make(chan string)
	go getChallenge(channel)
	challenge := <-channel
    fmt.Println(challenge)

	time.Sleep(time.Second * 3)
    fmt.Println("done")
}