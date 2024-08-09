package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // Creates a new map variable called m
	//Key     Value
	m["k1"] = 7            // Sets the value of the key "k1" to 7
	m["k2"] = 13           // Sets the value of the key "k2" to 13
	fmt.Println("map:", m) // Prints a string and the map

	v1 := m["k1"]          // Creates a new variable called v1 and sets its value to the value of the key (this time "k1") from the map m
	fmt.Println("v1:", v1) // Prints a string and the int variable v1

	v3 := m["k3"]          // Creates a new variable called v3 and sets its value to the value of the key (this time "k3") from the map m
	fmt.Println("v3:", v3) // Prints a string and the int variable v3

	fmt.Println("len:", len(m)) // Prints a string and the length of the map m

	delete(m, "k2")        // Deletes the Key (including its value) from the map
	fmt.Println("map:", m) // Prints a string and the map

	clear(m)               // Deletes everything inside the map
	fmt.Println("map:", m) // Prints a string and the map

	_, prs := m["k2"]        // Gets if the value of the key is set
	fmt.Println("prs:", prs) // Prints a string and the bool if the value is set

	//					Key    Value
	//					|      |
	n := map[string]int{"foo": 1, "bar": 2} // Creates a new map variable called n and sets its keys and values
	fmt.Println("map:", n)                  // Prints a string and the map

	n2 := map[string]int{"foo": 1, "bar": 2} // Creates a new map variable called n2 and sets its keys and values
	if maps.Equal(n, n2) {                   // Checks if map variable n and map variable m2 have the same keys with the same values
		fmt.Println("n == n2") // Prints a string
	}
}
