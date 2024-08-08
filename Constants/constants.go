package main

import (
	"fmt"  // Importing fmt (Format)
	"math" // Importing math for mathematical equations
)

const s string = "constant" // Creates constant string variable s and sets its value to "constant"

func main() {
	fmt.Println(s) // Prints string variable

	const n = 500000000 // Creates constant variable n and sets its value to 500000000

	const d = 3e20 / n       // Creates constant variable d and sets its value
	fmt.Println(d)           // Prints int variable d
	fmt.Println(int64(d))    // Prints int variable d as int64
	fmt.Println(math.Sin(n)) // Prints the sinus of const int variable n

	// A constant is set once and the variable cannot be changed
}
