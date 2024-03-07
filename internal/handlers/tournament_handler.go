package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rustem24liu/Golang-Final-Project/internal/tournament"
)

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	matches := tournament.RunTournament()
	jsonResponse, err := json.Marshal(matches)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
