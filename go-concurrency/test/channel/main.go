package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(500 * time.Millisecond)
	}
	close(c) // Closing the channel
}
