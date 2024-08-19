package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait() // Waits until every all Waitgroups from wg are Done.
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond) // Lets the function sleep for 500ms
	}
}
