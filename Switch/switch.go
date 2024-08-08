package main

import (
	"fmt"  // Importing fmt (Format)
	"time" // Importing time (for actual time information like the time right now or a day)
)

func main() {

	i := 2                         // Creates int variable i and sets it to value 2
	fmt.Print("Write ", i, " as ") // Prints a string without ending the line
	// Switches in go dont need a break; (like e.g. Java)
	switch i { // Switch for i
	case 1: // When i is 1 everything inside case 1: will be executed
		fmt.Println("one") // Prints a string
	case 2: // When i is 2 everything inside case 2: will be executed
		fmt.Println("two") // Prints a string
	case 3: // When i is 3 everything inside case 3: will be executed
		fmt.Println("three") // Prints a string
	}
	// Switches also have defaults with default: (like e.g. Java)
	switch time.Now().Weekday() { // Switch for every day in the Week (including weekend)
	case time.Saturday, time.Sunday: // When the actual time is Saturday or Sunday everything inside this case will be executed
		fmt.Println("It's the weekend") // Prints a string
	default: // When its not any case before this the default will be executed
		fmt.Println("It's a weekday") // Prints a string
	}

	t := time.Now()
	switch {
	case t.Hour() < 12: // When the actual time at the moment is below 12 am the case will be executed
		fmt.Println("It's before noon") // Prints a string
	default: // When its not any case before this the default will be executed
		fmt.Println("It's after noon") // Prints a string
	}

	whatAmI := func(i interface{}) { // Creates function whatAmI with an interface variable called i
		switch t := i.(type) { // Switches t and sets it to the type of i
		case bool: // When t is type of bool it executes the case
			fmt.Println("I'm a bool") // Prints a string
		case int: // When t is type of int it executes the case
			fmt.Println("I'm an int") // Prints a string
		default: // When t is an unknown type it executes the case
			fmt.Printf("Don't know type %T\n", t) // Prints a string with the variable type
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey") // Executes function whatAmI

}
