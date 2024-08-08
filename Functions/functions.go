package main

import "fmt" // Importing fmt (Format)

func plus(a int, b int) int { // A function that takes two parameters a & b both integers
	return a + b // Returns the value of a + b
}

func plusPlus(a, b, c int) int { // A function that takes three parameters a, b & c all integers
	return a + b + c // Returns the value of a + b +c
}

func main() {

	res := plus(1, 2)         // Creates a new int variable called res with the result of the function plus
	fmt.Println("1+2 =", res) // Prints a string with the int variable res

	res = plusPlus(1, 2, 3)     // Sets the int variable res with the result of the function plusPlus
	fmt.Println("1+2+3 =", res) // Prints a string with the int variable res
}
