package main

import "fmt" // Importing fmt (Format)

func main() {

	if 7%2 == 0 { // Checks if 7 is an even number
		fmt.Println("7 is even") // Prints when if is true
	} else { // When the if is not met it executes the following (when 7 is not even it does the following)
		fmt.Println("7 is odd") // Prints when if isnt true
	}

	if 8%4 == 0 { // Checks if 8 is divisible by 4
		fmt.Println("8 is divisible by 4") // Prints when if is true
	}

	if 8%2 == 0 || 7%2 == 0 { // Checks if either 8 or 7 is even
		fmt.Println("either 8 or 7 is even") // Prints when if is true
	}

	if num := 9; num < 0 { // Creates and sets num to value 9 and checks if num is lower smaller 0
		fmt.Println("is negative") // Prints when if is true
	} else if num < 10 { // Executes following when first if is false and when num is smaller than 10
		fmt.Println(num, "has 1 digit") // Prints when second if is true
	} else { // Executes following when both if's before are false
		fmt.Println(num, "has multiple digits") // Prints when both if's are false
	}
}
