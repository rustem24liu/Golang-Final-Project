package handlers

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

// ListPlayersHandler handles the redirection to list_of_players.html and fetches player data
func ListPlayersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database to fetch all players
	rows, err := db.Query("SELECT name, age, position FROM players")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold player data
	var players []Player

	// Iterate over the rows and populate the players slice
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.Name, &player.Age, &player.Position); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		players = append(players, player)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the list_of_players.html template with the player data
	tmpl, err := template.ParseFiles("/cmd/list_of_players/list_of_players.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, players)
}

// Player represents a player in the database
type Player struct {
	Name     string
	Age      int
	Position string
}
