package main

import (
	"net/http"
)

// resetHandler resets the fileserverHits counter to 0
func (cfg *apiConfig) handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Reset is only allowed in dev platform"))
		return
	}
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// reset the fileserverHits counter to 0
	cfg.fileserverHits.Store(0)
	cfg.databaseQueries.Reset(r.Context())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
