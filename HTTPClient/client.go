package main

import (
	"bufio"    // Import bufio (Scanner)
	"fmt"      // Import fmt (Format)
	"net/http" // Import HTTP (Web)
)

func main() {
	resp, err := http.Get("https://gobyexample.com") // Get Website
	if err != nil {                                  // Check for error with website connection
		panic(err) // Panic if error exists
	}
	defer resp.Body.Close() // If error close connection

	fmt.Println("Response status:", resp.Status) // Print Response Status

	scanner := bufio.NewScanner(resp.Body)     // Scanner for Website Body
	for i := 0; scanner.Scan() && i < 5; i++ { // Get first 5 lines of Website Body
		fmt.Println(scanner.Text()) // Print first 5 lines of Website Body
	}

	if err := scanner.Err(); err != nil { // Check for Error with Scanner
		panic(err) // Panic if error exists
	}
}
