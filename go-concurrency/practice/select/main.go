package main

import (
	"fmt"  // Imports the fmt-package (format)
	"time" // Imports the time-package
)

func main() {
	c1 := make(chan string) // Creates string channel 1 (c1)
	c2 := make(chan string) // Creates string channel 2 (c2)

	go func() { // Go-routine function (code continues and this runs next to it)
		for { // Repeats forever
			c1 <- "Every 500ms"                // Sends a message to channel 1 (c1)
			time.Sleep(time.Millisecond * 500) // Waits for 500 miliseconds before it continues (in this example repeats)
		}
	}()
	go func() { // Go-routine function (code continues and this runs next to it)
		for { // Repeats forever
			c2 <- "Every two seconds"   // Sends a message to channel 2 (c2)
			time.Sleep(time.Second * 2) // Waits for 2 seconds before it continues (in this example repeats)
		}
	}()

	for { // Repeats forever
		select { // Selects the channel
		case msg1 := <-c1: // When channel 1 (c1) has a message it will be send. It doesnt matter if channel 2 (c2) has a message.
			fmt.Println(msg1) // Prints Message 1 (msg1)
		case msg2 := <-c2: // When channel 2 (c2) has a message it will be send. It doesnt matter if channel 1 (c1) has a message.
			fmt.Println(msg2) // Prints Message 2 (msg2)
		}
	}
}
