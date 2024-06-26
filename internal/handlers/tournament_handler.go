package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/rustem24liu/Golang-Final-Project/internal/tournament"
)

func TournamentHandler(w http.ResponseWriter, r *http.Request) {
	result, err := tournament.RunTournament()
	if err != nil {
		http.Error(w, "Failed to generate tournament results", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("tournament/tournament.html")
	if err != nil {
		http.Error(w, "Failed to load HTML template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, "Failed to render HTML", http.StatusInternalServerError)
		return
	}
}
