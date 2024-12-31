package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	// user, err := cfg.db.CreateUser(r.Context(), params.Email)
	decoder := json.NewDecoder(r.Body)
	post := User{}
	if err := decoder.Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userEmail := post.Email
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// set http status code
	w.WriteHeader(http.StatusCreated)

	// set/write response
	newUser := User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     userEmail,
	}
	json.NewEncoder(w).Encode(newUser)
	return
}
