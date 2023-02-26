package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register the handler function to handle all requests
	http.HandleFunc("/", handler)

	// Start the server on port 8795
	fmt.Println("Listening on port 8795...")
	http.ListenAndServe(":8795", nil)
}