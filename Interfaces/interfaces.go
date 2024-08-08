package main

import (
	"fmt"  // Importing fmt (Format)
	"math" // Importing math (for mathematical equations)
)

// An interface is kind of like a struct but still different!
type geometry interface { // Creates an interface called geometry
	area() float64  // Takes a Function area() with the return type as float64
	perim() float64 // Takes a Function perim() with the return type as float64
}

type rect struct { // Creates a struct called circle
	width, height float64
}
type circle struct { // Creates a struct called circle
	radius float64
}

func (r rect) area() float64 { // A Function for the area of a rect. Returns float64 type
	return r.width * r.height // Returns the result of the equation
}
func (r rect) perim() float64 { // A Function for the perim of a rect. Returns float64 type
	return 2*r.width + 2*r.height // Returns the result of the equation
}

func (c circle) area() float64 { // A Function for the area of a circle. Returns float64 type
	return math.Pi * c.radius * c.radius // Returns the result of the equation
}
func (c circle) perim() float64 { // A Function for the perim of a circle. Returns float64 type
	return 2 * math.Pi * c.radius // Returns the result of the equation
}

func measure(g geometry) { // Function for measuring takes one parameter of type geometry
	fmt.Println(g)         // Prints the geometry
	fmt.Println(g.area())  // Prints the area of the geometry
	fmt.Println(g.perim()) // Prints the perim of the geometry
}

func main() {
	r := rect{width: 3, height: 4} // Creates a geometry variable called c and sets its type (rect), its width and height
	c := circle{radius: 5}         // Creates a geometry variable called c and kind of sets its type (circle) and its radius

	measure(r) // Measures the geometry from r
	measure(c) // Measures the geometry from c
}
