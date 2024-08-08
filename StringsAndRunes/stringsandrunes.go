package main

import (
	"fmt"          // Importing fmt (Format)
	"unicode/utf8" // Importing Unicode UTF8 for extra symbols
)

func main() {

	const s = "สวัสดี" // Creates a constant string variable called s

	fmt.Println("Len:", len(s)) // Prints a string with the length of the string variable s

	for i := 0; i < len(s); i++ { // Loops until i is not smaller than length of string variable s
		fmt.Printf("%x ", s[i]) // Prints unicode as string a doesnt end line
	}
	fmt.Println() // Prints an empty line

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Prints a string with the RuneCountInString from utf8

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // Prints a string with the runeValue and i which represents its position
	}

	fmt.Println("\nUsing DecodeRuneInString") // Prints a string
	for i, w := 0, 0; i < len(s); i += w {    // Loops until i isnt smaller than the length of string variable s
		runeValue, width := utf8.DecodeRuneInString(s[i:]) // Creates and sets runeValue and width using the DecodeRuneInString form utf8
		fmt.Printf("%#U starts at %d\n", runeValue, i)     // Prints a string with the runeValue and i which represents its position
		w = width                                          // Sets variable w to width

		examineRune(runeValue) // Executes function examineRune
	}
}

func examineRune(r rune) { // A function called examineRune which takes one parameter of the rune type

	if r == 't' { // Checks if rune variable r equals t
		fmt.Println("found tee") // Prints a string when if is true
	} else if r == 'ส' { // Checks if rune variable r equals ส
		fmt.Println("found so sua") // Prints a string when first if is false and this if is true
	}
}
