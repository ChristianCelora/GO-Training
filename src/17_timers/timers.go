package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	* Timers represent a single event in the future. 
	* You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. 
	* This timer will wait 2 seconds.
	*/
	timer1 := time.NewTimer(time.Second * 2)
	/**
	* The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	*/
	<-timer1.C
	fmt.Println("timer 1 fired")

	// timers can also be stopped before they fire
	timer2 := time.NewTimer(time.Second)
	go func () {
		<-timer2.C
		fmt.Println("timer 2 fired")
	} ()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped before firing")
	}

	/**
	* Tickers are for when you want to do something repeatedly at regular intervals.
	*/
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("tick, at:", t)
			case <-done:
				return
			}

		}
	} ()
	// stop ticker after 2 seconds (4 ticks)
	time.Sleep(time.Second * 2)
	ticker.Stop()	// stop ticker
	done <- true	// stop goroutine
}