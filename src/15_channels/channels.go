package main

import "fmt"

func sendMessage(channel chan string, message string) {
	channel <- message
}

func main() {
	/**
	* Channels are the pipes that connect concurrent goroutines. 
	* By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	*/
	channel := make(chan string)
	go sendMessage(channel, "hello")
	go sendMessage(channel, "world")
    fmt.Println(<-channel)
    fmt.Println(<-channel)

	/**
	* Buffered channels accept a limited number of values without a corresponding receiver for those values.
	*/

	buffered_channel := make(chan string, 2)
	/**
	* Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	* Notice: i am not using a goroutine for these functions. 
	* If i don't use it for an unbuffered channel i receive a "fatal error: all goroutines are asleep - deadlock!"
	*/
	sendMessage(buffered_channel, "hello")
	sendMessage(buffered_channel, "world")
	// These are syncronous, so i will recevie in order. The first print ha the possibility to start with "world"
    fmt.Println(<-buffered_channel)
    fmt.Println(<-buffered_channel)
}