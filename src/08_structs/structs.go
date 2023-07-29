package main
import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Christian", 26})
	fmt.Println(person{name: "Alice", age: 27})
	fmt.Println(person{name: "Micheal"})
	fmt.Println(person{})
	fmt.Println(&person{})	// pointer to struct. printed with "&" to indicate that it is a pointer

	bob := person{name: "Bob", age: 50}
	bobp := &bob
	fmt.Println(bob)
	fmt.Println(bobp)
	bob.age = bob.age + 1 
	fmt.Println(bobp)
}