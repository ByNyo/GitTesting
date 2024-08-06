package main

import (
	"fmt"
	"net/http" // Import HTTP (Web)
)

func hello(w http.ResponseWriter, req *http.Request) { // Hello Function for Hello Message
	fmt.Fprintf(w, "Hello\n") // Prints Hello Message
}

func headers(w http.ResponseWriter, req *http.Request) { // Headers Function for Testing
	for name, headers := range req.Header { // Repeats name with headers from Range of Requests Header
		for _, h := range headers { // Uses headers list to repeat each name
			fmt.Fprintf(w, "%v: %v\n", name, h) // Prints each Header
		}
	}
}

func main() { // Main Function where everything happens
	http.HandleFunc("/hello", hello)     // Using Hello Function to become Hello Message
	http.HandleFunc("/headers", headers) // Using Headers Function for Testing

	http.ListenAndServe(":8090", nil) // Listening to the Server under a Port (in this example port 8090)
}
