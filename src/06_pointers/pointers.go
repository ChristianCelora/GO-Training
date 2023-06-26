package main

import "fmt"

func increaseByOne(num *int) {
	*num = *num + 1 // *num access value
}

func main() {
	var i int
	i = 5
	fmt.Println("i, value:", i, ", address:", &i)
	increaseByOne(&i)
	fmt.Println("i, value:", i, ", address:", &i)
}