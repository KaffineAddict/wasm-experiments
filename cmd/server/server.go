package main

import (
	"fmt"
	"net/http"
)

// Q: Why?
// A: CORS - Cross Origin Resource Sharing
func main() {
	fmt.Println("serving")
	err := http.ListenAndServe(":9999", http.FileServer(http.Dir("bin")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
	fmt.Println("serving")
}
