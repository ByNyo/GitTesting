package main

import (
	"fmt"    // Importing fmt (Format)
	"slices" // Importing slices (e.g. for getting parts of an array)
)

func main() {
	var s []string                                    // Creates a new string array variable called s
	fmt.Println("uninit: ", s, s == nil, len(s) == 0) // Prints a string with information about the string array variable s

	s = make([]string, 3)                                  // Defines the string array variable s
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // Prints a string with added Information in between

	s[0] = "a"                // Sets the first string of the array variable s
	s[1] = "b"                // Sets the second string of the array variable s
	s[2] = "c"                // Sets the third string of the array variable s
	fmt.Println("set:", s)    // Prints a string and an array
	fmt.Println("get:", s[2]) // Prints a string and the string from the array variable s

	fmt.Println("len:", len(s)) // Prints a string and the length of an array variable s

	s = append(s, "d")      // Adds a string to the string array variable s
	s = append(s, "e", "f") // Adds a string to the string array variable s
	fmt.Println("apd:", s)  // Prints a string and the array variable s

	c := make([]string, len(s)) // Makes a new empty string array variable called c
	copy(c, s)                  // Copies the string array variable s into the other string array variable c
	fmt.Println("cpy:", c)      // Prints a string and the string array variable c

	l := s[2:5]           // Creates and gets/sets every string of the array starting from the second one and ending with the fifth one
	fmt.Println("sl1", l) // Prints a string and an array

	l = s[:5]             // Gets every string of the array starting ending with the fifth one
	fmt.Println("sl2", l) // Prints a string and an array

	l = s[2:]             // Gets every string of the array starting from the second one
	fmt.Println("sl3", l) // Prints a string and an array

	t := []string{"g", "h", "i"} // Makes a new string array
	fmt.Println("dcl:", t)       // Prints a string and an array

	t2 := []string{"g", "h", "i"} // Makes a new string array
	if slices.Equal(t, t2) {      // Checks if the "slices" (parts of the arrays) match
		fmt.Println("t == t2") // Prints a string
	}

	twoD := make([][]int, 3) // Makes a new int array and sets its length to 3
	for i := 0; i < 3; i++ { // Loops until i is not smaller than 3
		innerLen := i + 1               // Creates int variable innerLen and adds i + 1 each time the loop repeats
		twoD[i] = make([]int, innerLen) // Sets a new int Array in the i position of the string int array variable twoD
		for j := 0; j < innerLen; j++ { // Loops until j is not smaller than the int variable innerLen
			twoD[i][j] = i + j // Add i + j to both arrays of twoD
		}
	}
	fmt.Println("2d: ", twoD) // Prints a string with both arrays of twoD
}
