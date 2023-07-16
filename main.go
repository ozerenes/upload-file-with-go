package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Uploading File\n")

	// parse input
	r.ParseMultipartForm(10 << 20) // 10 MB limit

	// retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %v\n", handler.Filename)
	fmt.Printf("File Size: %v bytes\n", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)

	// write the temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	tempFile.Write(fileBytes)

	// return whether or not this has been successful
	fmt.Fprintf(w, "Successfully uploaded file\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go file upload")
	setupRoutes()
}
