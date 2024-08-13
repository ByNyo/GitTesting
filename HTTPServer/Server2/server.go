package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/helloworld", hello)

	http.ListenAndServe(":9000", nil)
}
