package main

import (
	"fmt"
	"net/http" // Import HTTP (Web)
)

func helloworld(w http.ResponseWriter, req *http.Request) { // Hello Function for Hello Message
	fmt.Fprintf(w, "Hello, World!\n") // Prints Hello Message
}

func main() { // Main Function where everything happens
	http.HandleFunc("/helloworld", helloworld) // Using helloworld Function to become Hello Message

	http.ListenAndServe(":8091", nil) // Listening to the Server under a Port (in this example port 8092)
}
