package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}     // Creates an int array variable called nums and sets its values
	sum := 0                   // Creates an int variable
	for _, num := range nums { // Loops until its out of range of nums and dumps the index
		sum += num // Adds num to the int variable sum each loop
	}
	fmt.Println("sum:", sum) // Prints a string and int variable sum

	for i, num := range nums { // Loops until its out of range of nums
		if num == 3 { // Checks if the num is 3
			fmt.Println("index:", i) // Prints a string along with the int variable i
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // Creates a map variable and sets its keys and values
	for k, v := range kvs {                               // Loops until its out of range of the map variable kvs and separates the key from the value
		fmt.Printf("%s -> %s\n", k, v) // Prints a string and the key along with the value
	}

	for k := range kvs { // Loops until its out of range of the map variable kvs, gets the key and dumps the value
		fmt.Println("key:", k) // Prints a string and the key
	}

	for i, c := range "go" { // Loops over each character in the string "go"
		fmt.Println(i, c) // Prints the index (i) and the unicode point (c) of the character
	}
}
