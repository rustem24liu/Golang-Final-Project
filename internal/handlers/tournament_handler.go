package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rustem24liu/Golang-Final-Project/internal/tournament"
)

func TournamentHandler(w http.ResponseWriter, r *http.Request) {
	// Run the tournament
	result := tournament.RunTournament()

	// Convert the result to JSON
	jsonResult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal tournament result", http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResult)
}
