package main

import (
	"encoding/json"
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// load the current value of the fileserverHits counter
	hits := cfg.fileserverHits.Load()
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	// write status code in the response
	w.WriteHeader(http.StatusOK)
	// write the current value of the fileserverHits counter in the response
	htmlContent := fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>
	`, hits)

	w.Write([]byte(htmlContent))
}

// resetHandler resets the fileserverHits counter to 0
func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// reset the fileserverHits counter to 0
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type errorResponse struct {
	Error string `json:"error"`
}

func chirpHandler(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// create struct to store the post
	type chirp struct {
		Body string `json:"body"`
	}
	// decode the request body
	decoder := json.NewDecoder(r.Body)
	// initialize a post struct
	post := chirp{}
	// decode the request body into the post struct
	// the post struct will store the post in post.Body if no error occurs
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(post.Body) > 140 {
		// write the status code 400 in the response
		w.WriteHeader(http.StatusBadRequest)
		// write the error message in the response
		// errorResponse struct is used to format the error message in JSON
		json.NewEncoder(w).Encode(errorResponse{Error: "Chirp is too long"})
		return
	}
	type validResponse struct {
		Valid bool `json:"valid"`
	}
	fmt.Println(post.Body)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(validResponse{Valid: true})
	return
}

func main() {

	mux := http.NewServeMux()

	apiCfg := &apiConfig{}

	fileServer := http.FileServer(http.Dir("."))

	// wrap the file server with the middlewareMetricsInc middleware
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	mux.HandleFunc("POST /admin/reset", apiCfg.resetHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("POST /api/validate_chirp", chirpHandler)
	mux.HandleFunc("GET /api/healthz", readinessHandler)

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
