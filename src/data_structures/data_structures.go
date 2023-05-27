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
	* 
	* Internally slices composed by 3 parts:
	* - a pointer to the subsequent array
	* - a integer to indicate the current length of the slice
	* - an integer to indicate the maxium capacity of the slice
	*/
	var empty_slice []int
	fmt.Println("empty_slice", empty_slice)
	
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("slice", slice)
	/**
	* Slicing a slice will create a new slice 
	* with the pointer pointing to the lower bound (if specified) and updated the length and capacity
	* 
	* 	new = slice[1:] (len: 3, cap: 3)
	* 
	*			[0]
	* 	new -> 	[1]
	*			[2]
	*			[3]
	*/
	fmt.Println("slice sliced", slice[1:]) 
	/**
	* A slice can move the upper bound hight to increase his length, but not lower the lower bound
	* [0, 1, 2, 3]
	* slice[:2] 
	* slice -> [0, 1]
	* len: 2
	* cap: 4
	*/
	fmt.Println("slice sliced", slice[:2]) // upper index not included

	/**
	* After a slice has been sliced all the indexed are updated
	*/
	smaller_slice := slice[1:]
	// smaller_slice[0] === slice[1]
	// smaller_slice[1] === slice[2]
	fmt.Println("slice[0]", slice[0], "smaller_slice[0]", smaller_slice[0]) 
	fmt.Println("slice[0]", slice[1], "smaller_slice[1]", smaller_slice[1]) 
	fmt.Println("slice[0]", slice[2]) 

	/**
	* During an append a new array is created with len(slice) and the pointer is changed to the new array
	* The old array will be managed by the garbage collector (if not used)
	*
	* 			[1, 2, 3]
	* slice ->	[1, 2, 3, 4, 5, 6]
	* len = 6
	* cap = 6
	*/
	slice = append(slice, 4, 5, 6)
	fmt.Println("slice", slice)

	var s1 []int
	s1 = append(s1, 1, 2, 3)
	fmt.Println("slice", s1, "l", len(s1), "c", cap(s1))

	// Map
	var my_map = make(map[int]string)
	my_map[1] = "one"
	my_map[2] = "two"
	my_map[3] = "three"

	fmt.Println("map", my_map)

	delete(my_map, 2)

	fmt.Println("map", my_map)

	/**
	* If the key doesn’t exist, the zero value of the value type is returned.
	* Zero values list:
	* boolean - false
	* integer - 0
	* string - ""
	* pointer - nil
	*/
	fmt.Println(my_map[4])

	/**
	* The optional second return value when getting a value from a map indicates if the key was present in the map. 
	* This can be used to disambiguate between missing keys and keys with zero values like 0 or ""
	*/
	_, prs1 := my_map[1]
	fmt.Println("existing key", prs1)

	_, prs2 := my_map[4]
	fmt.Println("non-existing key", prs2)

	// Channel
}