package main

import "fmt"

type person struct {
	name string
	age uint32
	phone
}

type phone struct {
	label string
	number string
	prefix string
}

func (p phone) getNumberWithPrefix() string {
	return "+" + p.prefix + p.number
}

func main() {
	// When creating structs with literals, we have to initialize the embedding explicitly
	// Here the embedded type serves as the field name.
	bob := person{
		name: "Bob",
		age: 32,
		phone: phone{
			label: "Work",
			number: "33932456789",
			prefix: "39",
		},
	}

	fmt.Println(bob)
	// the fields and methods of the embedded struct can be accessed from the main struct
	fmt.Println("phone:", bob.number, bob.phone.number)
	fmt.Println("phone with prefix:", bob.getNumberWithPrefix(), bob.phone.getNumberWithPrefix())
}
