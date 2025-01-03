package main

import (
	"encoding/json"
	"example5/internal/database"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	User_ID   uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
}

func filterWords(originalString string, badWords map[string]int) string {
	lowercaseText := strings.ToLower(originalString)
	lowercaseWords := strings.Split(lowercaseText, " ")
	filteredWords := strings.Split(originalString, " ")
	for i, lowercaseWord := range lowercaseWords {
		if _, exists := badWords[lowercaseWord]; exists {
			filteredWords[i] = "****"
		}
	}
	filteredText := strings.Join(filteredWords, " ")
	return filteredText
}

func (cfg *apiConfig) handleCreateChirps(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// creates new JSON decoder that will read from the request body
	decoder := json.NewDecoder(r.Body)
	// initialize a post struct
	post := Chirp{}
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

	badWords := map[string]int{
		"kerfuffle": 0,
		"sharbert":  0,
		"fornax":    0,
	}

	filteredText := filterWords(post.Body, badWords)
	w.WriteHeader(http.StatusCreated)
	// get user_id from the request and create a new uuid
	params := database.CreateChirpParams{Body: filteredText, UserID: post.User_ID}

	newChirp, err := cfg.databaseQueries.CreateChirp(r.Context(), params)
	if err != nil {
		log.Printf("error creating chirp: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// encodes `Chirp` struct as json and writes it to the http response `w`
	json.NewEncoder(w).Encode(newChirp)
	return
}

// set headers
// write http status code on response
// create response variable that will be encoded into json
// encode response into json
func (cfg *apiConfig) handleChirpsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	chirps := make([]Chirp, 10)
	// get chirps from postgresql
	json.NewEncoder(w).Encode(chirps)
	return
}
