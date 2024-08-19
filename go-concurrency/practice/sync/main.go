package main

import (
	"fmt"  // Imports the fmt-package (format)
	"sync" // Imports the sync-package
	"time" // Imports the time-package
)

func main() {
	var wg sync.WaitGroup // Creates a WaitGroup from the sync-package
	wg.Add(1)             // Adds 1 to wg

	go func() { // Go-routine function (code continues and this runs next to it)
		count("sheep") // Executes the count function
		wg.Done()      // Sets wg to Done
	}()

	wg.Wait() // Waits until wg is Done
}

func count(thing string) { // Count function that takes two parameter, a string and a string channel
	for i := 1; i <= 5; i++ { // Repeats until i is not smaller or equal to 5
		fmt.Println(i, thing)              // Prints the int variable i and the string variable thing
		time.Sleep(time.Millisecond * 500) // Waits for 500 miliseconds before it continues (in this example repeats or continues)
	}
}
