package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try /helloworld")
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func uploadfile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := r.FormFile("myfile")

	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	// Upload File
	tempFile, err2 := ioutil.TempFile("uploads", "upload-*.txt")
	if err2 != nil {
		fmt.Println("Err2:", err2)
	}
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println("Err3:", err3)
	}
	tempFile.Write(fileBytes)
	fmt.Println("Done")
}

func handleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/helloworld", helloworld)
	http.HandleFunc("/upload", uploadfile)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
