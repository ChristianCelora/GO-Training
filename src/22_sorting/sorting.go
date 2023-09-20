package main

import (
	"fmt"
	"cmp"
	"slices"
)

type Person struct {
	name string
	age int
}

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

	// We can sort by a custom function using slices.SortFunc
	// cmp.Compare is helpful for this
	age_sort := func (a, b Person) int {
		return cmp.Compare(a.age, b.age)
	}
	persons := []Person{
		Person{name: "aldo", age: 22},
		Person{name: "mario", age: 18},
		Person{name: "giovanni", age: 33},
		Person{name: "maria", age: 25},
	}
	slices.SortFunc(persons, age_sort)
	fmt.Println("Contacts sorted", persons)
}
