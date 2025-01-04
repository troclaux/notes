package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleChirpGet(w http.ResponseWriter, r *http.Request) {

	// get chirp id from the URL path
	strID := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(strID)
	if err != nil {
		log.Println("Error parsing UUID string:", err)
		return
	}

	// attempt to get the chirp from the database
	// IMPORTANT: convert the dbChirp to a Chirp struct
	dbChirp, err := cfg.databaseQueries.GetChirp(r.Context(), chirpID)
	if err != nil {
		log.Printf("error getting chirp: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	chirp := Chirp{
		ID:        dbChirp.ID,
		UserID:    dbChirp.UserID,
		Body:      dbChirp.Body,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
	}

	// write the chirp to the response
	if err := json.NewEncoder(w).Encode(chirp); err != nil {
		log.Printf("error encoding response: %v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
