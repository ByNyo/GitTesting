package main

import "fmt" // Importing fmt (Format)

func main() {

	var a = "initial" // Creates a new string variable called a and sets it to "initial"
	fmt.Println(a)    // Prints string variable a

	var b, c int = 1, 2 // Creates 2 new integer variables and sets their values
	fmt.Println(b, c)   // Prints 2 integer variables b & c

	var d = true   // Creates a new boolean variable called d and sets it to true
	fmt.Println(d) // Prints boolean variable d

	var e int      // Creates a new integer variable called e, its value is default which is 0
	fmt.Println(e) // Prints integer variable e

	f := "apple"   // Creates a new string variable called f and sets its value to "apple"
	fmt.Println(f) // Prints string variable f
}
