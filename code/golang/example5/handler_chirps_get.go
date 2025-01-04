package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handleChirpsGet(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.databaseQueries.GetChirps(r.Context())
	if err != nil {
		log.Printf("error getting chirps: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// attempt to get the chirp from the database
	// IMPORTANT: convert the dbChirps to a map of Chirps{}
	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	if err := json.NewEncoder(w).Encode(chirps); err != nil {
		log.Printf("error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
