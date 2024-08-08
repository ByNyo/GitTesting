package main

import "fmt" // Importing fmt (Format)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals() // Gets 2 return values from function vals and puts them into different int variables
	fmt.Println(a) // Prints the int variable a
	fmt.Println(b) // Prints the int variable b

	_, c := vals() // Gets 2 return values from function vals. Then dumps the first return and puts the second into an int variable called c
	fmt.Println(c) // Prints the int variable c
}
