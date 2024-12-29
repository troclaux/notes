package main

import (
	"database/sql"
	"encoding/json"
	"example5/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	// atomic.Int32 is a type that provides atomic operations for int32 values
	fileserverHits  atomic.Int32
	databaseQueries *database.Queries
	platform        string
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
func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, r *http.Request) {
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
func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// reset the fileserverHits counter to 0
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
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

type chirp struct {
	Body string `json:"body"`
}

func handleChirp(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// creates new JSON decoder that will read from the request body
	decoder := json.NewDecoder(r.Body)
	// initialize a post struct
	post := chirp{}
	// attempts to decode json from the request body and stores data into the struct
	if err := decoder.Decode(&post); err != nil {
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

	lowercaseText := strings.ToLower(post.Body)
	lowercaseWords := strings.Split(lowercaseText, " ")
	filteredWords := strings.Split(post.Body, " ")
	for i, lowercaseWord := range lowercaseWords {
		if lowercaseWord == "kerfuffle" || lowercaseWord == "sharbert" || lowercaseWord == "fornax" {
			filteredWords[i] = "****"
			continue
		}
	}
	filteredText := strings.Join(filteredWords, " ")

	type cleaned_response struct {
		Cleaned_body string `json:"cleaned_body"`
	}
	fmt.Println(filteredText)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cleaned_response{Cleaned_body: filteredText})
	return
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// establishes a connection pool to the database that manages multiple connections
	// the connection string is stored in the DB_URL environment variable
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		println("Error connecting to database")
	}
	defer db.Close()

	dbQueries := database.New(db)

	// creates new http request multiplexer
	mux := http.NewServeMux()

	// instantiate struct with request counter and connection pool
	apiCfg := &apiConfig{databaseQueries: dbQueries, platform: os.Getenv("PLATFORM")}

	// serves files to the client from the defined path
	fileServer := http.FileServer(http.Dir("."))

	// wrap the file server with the middlewareMetricsInc middleware
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	mux.HandleFunc("POST /admin/reset", apiCfg.handleReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handleMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handleChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handleUser)
	mux.HandleFunc("GET /api/healthz", handleReadiness)

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
