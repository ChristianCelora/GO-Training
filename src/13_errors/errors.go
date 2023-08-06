package main

import (
    "fmt"
	"errors"
)

type argError struct {
	arg float32
	msg string
}

// A custom error should implement an Error() function which returns the message string
func (e *argError) Error() string {
	return fmt.Sprintf("%s. value: %2.f", e.msg, e.arg)
}

func division(a float32, b float32) (float32, error) {
	if b == 0 {
		return -1, &argError{b, "cannot divide by zero"}
	}

	return a / b, nil
}

func willFailForNumber42(a int) (int, error) {
	if a == 42 {
		// errors.New constructs a basic error value with the given error message
		return a, errors.New("Failed. Told ya")
	}

	return a, nil
}

func main() {
	var res float32
	var err error

	res, err = division(3, 5)
	fmt.Println(res, err)

	res, err = division(7, 0)
	fmt.Println(res, err)

	fmt.Println(willFailForNumber42(41))
	fmt.Println(willFailForNumber42(42))
	fmt.Println(willFailForNumber42(43))
}