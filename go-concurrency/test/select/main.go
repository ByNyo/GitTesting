package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // Creates a string channel variable called c1
	c2 := make(chan string) // Creates a string channel variable called c2
	go func() {             // goroutine
		for { // Repeats forever
			c1 <- "Every 500ms"                // Sets the string of the channel c1
			time.Sleep(500 * time.Millisecond) // Lets the goroutine sleep for 500 miliseconds
		}
	}()
	go func() { // goroutine
		for { // Repeats forever
			c1 <- "Every 2 seconds"     // Sets the string of the channel c2
			time.Sleep(2 * time.Second) // Lets the goroutine sleep for 2 seconds
		}
	}()
	for { // Repeats forever
		select { // Selects any of the cases below
		case msg1 := <-c1: // <- before a channel gives the output for example a string to for example a variable.
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
