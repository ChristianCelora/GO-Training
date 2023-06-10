package main

import "fmt"

func main() {
	var a int = 50
	var b string = "050"
	fmt.Println(a, b)

	var a2 int = 050 // this is octal (40 in decimal)
	var a3 int = 0x50 // this is hexadecimal
	fmt.Println(a2, a3)

	const c = "cannot change me"
	// c = "compiler error" // compiler error
	fmt.Println(c)

	d := 40
	e := true
	f := "hey"
	fmt.Println(d, e, f)

	var g, h int = 2, 3
	var i, j int = 4, 5
	fmt.Println(g, h)
	fmt.Println(i, j)

	// if a variable is not used is the compiler return an error
	// k := 5

	// operations

	// basic math
	a2 = 3
	a3 = 4
	l1 := a2 + a3
	l2 := a2 - a3
	l3 := a2 * a3
	l4 := a2 / a3
	l5 := a2 % a3
	l6 := -a2
	l7 := +a2
	fmt.Println(l1, l2, l3, l4, l5, l6, l7)
}