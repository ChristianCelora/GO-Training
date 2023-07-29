package main

import "fmt"

// Go supports methods defined on struct types.
type rectangle struct {
	height, width float32
}

// This area method has a receiver type of *rect.
func (r *rectangle) area() float32 {
	return r.height * r.width
}

func (r *rectangle) perim() float32 {
	return 2*r.height + 2*r.width
}

func main() {
	r := rectangle{height: 42.90, width: 33}
	area := r.area() 
	perim := r.perim()
	fmt.Println("Rectangle", r, "area:", area, "perim:", perim)
	// area is 1415.7001. Golang has some flaoting point issues with math operations
	// One solution is to show or round the variable to the needed deciaml places
	// If you're dealing with currencies it can be tricky. We could store them in a struct with dollars and cents
	fmt.Printf("area: %.2f \n", area)
}