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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(chirps); err != nil {
		log.Printf("error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
