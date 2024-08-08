package main

import "fmt"

type person struct { // Creates a new struct called person
	name string // A string that can be set in a variable with the person type
	age  int    // An int that can be set in a variable with the person type
}

func newPerson(name string) *person { // A function that takes a string and return type of person
	p := person{name: name} // Creates person with name
	p.age = 42              // Sets the persons age to 42
	return &p               // Returns person type
}

func main() {
	fmt.Println(person{"Bob", 20})              // Prints a person
	fmt.Println(person{name: "Alice", age: 30}) // Prints a person
	fmt.Println(person{name: "Fred"})           // Prints a person
	fmt.Println(&person{name: "Ann", age: 40})  // Prints a person
	fmt.Println(newPerson("Jon"))               // Prints a person
	s := person{name: "Sean", age: 50}          // Prints a person
	fmt.Println(s.name)                         // Prints a name

	sp := &s            // Gets a person from the person variable s
	fmt.Println(sp.age) // Prints an age

	sp.age = 51         // Sets the age of person variable sp
	fmt.Println(sp.age) // Prints an age

	dog := struct { // Creates a new struct called dog
		name   string // A string that can be set in a variable with the dog type
		isGood bool   // A bool that can be set in a variable with the dog type
	}{ // Creates a dog
		"Rex", // Name of the dog
		true,  // Bool if the dog is good
	}
	fmt.Println(dog) // Prints the dog
}
