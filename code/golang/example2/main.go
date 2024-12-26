package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a server multiplexer, which is a request router
	// Decides which function to call based on the URL path
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	fmt.Println("Server is running on http://localhost:8080")

	// ListenAndServe starts an HTTP server with a given address and handler
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
