package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	// atomic.Int32 is a type that provides atomic operations for int32 values
	fileserverHits atomic.Int32
}

// middlewareMetricsInc increments the fileserverHits counter for each request
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// increment the fileserverHits counter
		cfg.fileserverHits.Add(1)
		// after incrementing the counter, the request is passed to the next handler
		next.ServeHTTP(w, r)
	})
}

// metricsHandler returns the current value of the fileserverHits counter
func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	// load the current value of the fileserverHits counter
	hits := cfg.fileserverHits.Load()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// write status code in the response
	w.WriteHeader(http.StatusOK)
	// write the current value of the fileserverHits counter in the response
	w.Write([]byte(fmt.Sprintf("Hits: %d", hits)))
}

// resetHandler resets the fileserverHits counter to 0
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	// reset the fileserverHits counter to 0
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func main() {

	mux := http.NewServeMux()

	apiCfg := &apiConfig{}

	fileServer := http.FileServer(http.Dir("."))

	// wrap the file server with the middlewareMetricsInc middleware
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))
	mux.HandleFunc("/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("/reset", apiCfg.resetHandler)

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
