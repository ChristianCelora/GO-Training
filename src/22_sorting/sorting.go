package main

import (
	"fmt"
	"slices"
)

/**
 * Go’s slices package implements sorting for builtins and user-defined types
 * From go 1.21
 * Documentation: https://pkg.go.dev/slices
*/
func main() {
	// Sorting functions are generic, and work for any ordered built-in type
	names := []string{"mario", "luca", "alberto", "giovanni"}
	fmt.Println("sorted", slices.IsSorted(names), "|", names) // IsSorted returns if a slice is sorted
	slices.Sort(names)
	fmt.Println("sorted", slices.IsSorted(names), "|", names)

	numbers := []int{1, 4, 7, 10, 3, 5, 5}
	fmt.Println("sorted", slices.IsSorted(numbers), "|", numbers)
	slices.Sort(numbers)
	fmt.Println("sorted", slices.IsSorted(numbers), "|", numbers)
}
