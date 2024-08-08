package main

import "fmt" // Importing fmt (Format)

func main() {

	i := 1       // Creates int variable i and sets its value to 1
	for i <= 3 { // Loops until int variable i is not smaller or equal than 3
		fmt.Println(i) // Prints int variable i each time
		i = i + 1      // Adds one to int variable i each time
	}

	for j := 0; j < 3; j++ { // Creates int variable j and sets it to 0. Repeats/Loops until j is not smaller than 3 and adds 1 each time it repeats/loops
		fmt.Println(j) // Prints int variable j
	}

	for i := range 3 { // Loops until int variable i is in range of 3 (i == 3)
		fmt.Println("range", i) // Prints "range" and the value of int variable i
	}

	for { // Loops without anything neccessary
		fmt.Println("loop") // Prints string "loop"
		break               // Breaks the loop
	}

	for n := range 6 { // Loops until int variable n is in range of 6 (n == 6)
		if n%2 == 0 { // Checks if n%2 equals 0
			continue // Continues when if is met
		}
		fmt.Println(n) // Prints int variable n
	}
}
