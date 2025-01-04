package main

import (
	"database/sql"
	"example5/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
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

func main() {

	// load env vars from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// establishes a connection pool to the database that manages multiple connections
	// the connection string is stored in the DB_URL environment variable
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()

	dbQueries := database.New(db)

	// creates new http request multiplexer
	mux := http.NewServeMux()

	// initialize struct with request counter and connection pool
	apiCfg := &apiConfig{databaseQueries: dbQueries, platform: os.Getenv("PLATFORM")}

	// serves files to the client from the defined path
	fileServer := http.FileServer(http.Dir("."))

	// wrap the file server with the middlewareMetricsInc middleware
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	mux.HandleFunc("POST /admin/reset", apiCfg.handleReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handleMetrics)
	mux.HandleFunc("POST /api/users", apiCfg.handleUsersCreate)
	mux.HandleFunc("GET /api/healthz", handleReadiness)
	mux.HandleFunc("POST /api/chirps", apiCfg.handleCreateChirps)
	mux.HandleFunc("GET /api/chirps", apiCfg.handleChirpsGet)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handleChirpGet)

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
