package main

import (
	"os"
	"errors"
)

func main() {
	_, err := os.Create("/tmp/test")
	err = errors.New("Error while creating the file") // test
	if err != nil {
		/**
		 * A panic typically means something went unexpectedly wrong
		 * Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status
		*/
		panic(err)
	}
}