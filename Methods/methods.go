package main

import "fmt" // Importing fmt (Format)

type rect struct {
	width, height int // Two ints that can be set in a variable with the rect type
}

func (r *rect) area() int { // Function called area which is for rect types and returns an int
	return r.width * r.height // Returns the result of the equation
}

func (r rect) perim() int { // Function called perim which is for rect types and returns an int
	return 2*r.width + 2*r.height // Returns the result of the equation
}

func main() {
	r := rect{width: 10, height: 5}  // Creates a rect variable called r and sets its values
	fmt.Println("area: ", r.area())  // Prints a string and the area of the rect variable r
	fmt.Println("perim:", r.perim()) // Prints a string and the perim of the rect variable r

	rp := &r                          // Creates rect variable rp and gets its values from r
	fmt.Println("area: ", rp.area())  // Prints a string and the area of the rect variable rp
	fmt.Println("perim:", rp.perim()) // Prints a string and the perim of the rect variable rp
}
