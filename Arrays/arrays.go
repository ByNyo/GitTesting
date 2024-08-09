package main

import "fmt" // Importing fmt (Format)

func main() {
	var a [5]int           // Creates an int array with 5 values called a
	fmt.Println("emp:", a) // Prints a string and the array

	a[4] = 100                // Sets the fourth value of the array a
	fmt.Println("set:", a)    // Prints a string and the edited array
	fmt.Println("get:", a[4]) // Prints a string and the fourth value of the array

	fmt.Println("len:", len(a)) // Prints a string and length of the array

	b := [5]int{1, 2, 3, 4, 5} // Creates an int array with 5 values called 5 it sets the values in order 1,2,3,4,5
	fmt.Println("dcl:", b)     // Prints a string and the array

	b = [...]int{1, 2, 3, 4, 5} // Sets/Changes the values of the array b
	fmt.Println("dcl:", b)      // Prints a string and the array

	b = [...]int{100, 3: 400, 500} // Sets/Changes the values of the array b
	fmt.Println("idx:", b)         // Prints a string and the array

	var twoD [2][3]int       // Creates two int arrays, one with 2 values and 3 values called twoD
	for i := 0; i < 2; i++ { // Loops until i is not smaller than 2
		for j := 0; j < 3; j++ { // Loops until j is not smaller than 3
			twoD[i][j] = i + j // Add i + j to both arrays of twoD
		}
	}
	fmt.Println("2d: ", twoD) // Prints a string with both arrays

	twoD = [2][3]int{ // Changes the values of both arrays from twoD
		{1, 2, 3}, // Values changed for the first array
		{1, 2, 3}, // Values changed for the second array
	}
	fmt.Println("2d: ", twoD) // Prints a string with both arrays of twoD
}
