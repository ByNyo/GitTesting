package main

import "fmt" // Importing

func sum(nums ...int) { // A function that takes an unlimited number of ints
	fmt.Print(nums, " ") // Prints int variable nums and a string
	total := 0           // Creates a new int variable called total

	for _, num := range nums { // Creates int variable num from nums and loops until it is no longer in range of nums
		total += num // Adds int variable num to the int variable total
	}
	fmt.Println(total) // Prints the int variable total
}

func main() {
	sum(1, 2)    // Executes the function sum
	sum(1, 2, 3) // Executes the function sum

	nums := []int{1, 2, 3, 4} // Creates a int array variable called nums
	sum(nums...)              // Executes the function sum with the int array variable nums
}
