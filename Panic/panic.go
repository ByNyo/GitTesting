package main

import "os"

func main() {

	panic("a problem") // panic with custom text

	_, err := os.Create("/tmp/file") // Error made on Purpose for testing
	if err != nil {                  // if error exists
		panic(err) // panic (writes out the error)
	}
}
