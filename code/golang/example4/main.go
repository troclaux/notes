package main

import (
	"fmt"
	"net/http"
)

// readinessHandler returns a simple OK response
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/", fileServer)
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))
	mux.Handle("/app/assets/", http.StripPrefix("/app", fileServer))
	mux.HandleFunc("/healthz", readinessHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
