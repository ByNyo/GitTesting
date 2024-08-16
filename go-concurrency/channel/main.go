package main

import (
	"fmt"  // Imports the fmt-package (format)
	"time" // Imports the time-package
)

func main() {
	c := make(chan string) // Creates a string channel (c)

	go count("sheep", c) // Calls the function count() as a go-routine so the code afterwards continues.
	for msg := range c { // For every message it will repeat whats inside
		fmt.Println(msg) // Prints a messages from the channel
	}
}

func count(thing string, c chan string) { // Count function that takes two parameter, a string and a string channel
	for i := 1; i <= 5; i++ { // Repeats until i is not smaller or equal to 5
		c <- thing                         // Sends string variable thing to the channel c
		time.Sleep(time.Millisecond * 500) // Waits for 500 miliseconds before it continues (in this example repeats or continues)
	}
	close(c)
}
