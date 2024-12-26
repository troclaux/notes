package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Handle registers the handler function for the given pattern
	mux.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
