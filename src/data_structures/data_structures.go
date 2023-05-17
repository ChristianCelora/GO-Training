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
	* "make" in Golang is a keyword that is used to initialize only slices, maps, and channels.
	* It does not return the pointer to the item. 
	* It initializes the data in memory with zero value
	*/

	/**
	* Slice
	* Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 
	* To create an empty slice with non-zero length, use the builtin make
	*/
	var empty_slice []int
	fmt.Println("empty_slice", empty_slice)
	
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("slice", slice)
	fmt.Println("slice sliced", slice[1:]) 
	fmt.Println("slice sliced", slice[:2]) // upper index not included

	slice = append(slice, 4, 5, 6)
	fmt.Println("slice", slice)

	var s1 []int
	s1 = append(s1, 1, 2, 3)
	fmt.Println("slice", s1, "l", len(s1), "c", cap(s1))

	// Map

	// Channel
}