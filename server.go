package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening on port 8795...")
	http.ListenAndServe(":8795", nil)
}
