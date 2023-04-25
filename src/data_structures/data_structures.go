package main

import "fmt"

func main(){
	// Array
	var arr [5]int
	fmt.Println("arr of size", len(arr), ":", arr)
	arr[1] = 15
	fmt.Println("arr", arr)
	// arr[10] = 12 // compile error: index out of bounds

	// create and init in same line
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2", arr2)
	for i := 1; i < len(arr2); i++ {
        arr2[i] += arr2[i-1]
    }
	fmt.Println("arr2", arr2)

	// b is a new Array with the values of a
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := a
	a[1] = 0
	fmt.Println("Array a:", a)
	fmt.Println("Array b:", b)


	/**
	* Slice
	* Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 
	* To create an empty slice with non-zero length, use the builtin make
	*/
	// slice := make([]int)

	// Map

	// Channel
}