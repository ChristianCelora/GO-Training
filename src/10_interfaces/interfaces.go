package main

import (
	"fmt"
	"math"
	"reflect"
)

// Interfaces are named collections of method signatures.
type geometry interface {
    area() float32
    perim() float32
}

type rectangle struct {
    width, height float32
}

// Notice: there is no pointer in the receiver
func (r rectangle) area() float32 {
	return r.height * r.width
}

func (r rectangle) perim() float32 {
	return 2*r.height + 2*r.width
}

type circle struct {
    radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

type dot struct {
    radius float32
}

func printMeasurements(g geometry) {
    fmt.Println(reflect.TypeOf(g))	// get struct
    fmt.Println(g)
    fmt.Println(g.area())
    // fmt.Println(g.perim())
}

func main() {
	r := rectangle{width: 10.24, height: 32.56}
	c := circle{radius: 15.4}

	printMeasurements(r)
	printMeasurements(c)

	// The next line will throw an error during compile
	// In golang we do not specify the interface in the struct, but when it is used an interface,
	// it checks that the struct has implemented ALL the methods of the interface
	// d := dot{radius: 12.3}
	// printMeasurements(d) 
}