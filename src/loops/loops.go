package main

import "fmt"

func main(){
	// for / while 
	// use FOR for everything 
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("for")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// we use for as while loops
	sum := 0
	fmt.Println("while")
	for sum < 6 {
		sum += 1
		fmt.Println(sum)
	}

	// break / continue
	for i := 0; i < len(a); i++ {
		if a[i] > 2 {
			fmt.Println("break at index", i)
			break
		}

	}
	fmt.Println("still a, but only odds")
	for i := 0; i < len(a); i++ {
		if a[i] % 2 == 0 {
			continue
		}
		fmt.Println(a[i])
	}
}