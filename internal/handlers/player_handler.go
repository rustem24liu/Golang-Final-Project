package handlers

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "log"
	"net/http"
	_ "net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rustem24liu/Golang-Final-Project/internal/repository"
	"github.com/rustem24liu/Golang-Final-Project/models"
)

type PlayerHandler struct {
	playerRepo *repository.PlayerRepo
}

func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	return &PlayerHandler{
		playerRepo: repository.NewPlayerRepo(db),
	}
}

func (ph *PlayerHandler) ListOfAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.ListOfAllPlayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination, sorting, and filtering
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	sortBy := r.URL.Query().Get("sort")
	// Example filtering parameter, you can add more as needed
	positionFilter := r.URL.Query().Get("player_pos")
	fmt.Println(positionFilter, "this is position filter")

	// Construct filter map
	filters := make(map[string]interface{})
	if positionFilter != "" {
		filters["player_pos"] = positionFilter
	}

	// Retrieve players from repository
	players, err := ph.playerRepo.GetAllPlayers(pageNum, pageSize, sortBy, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error retrieving players:", err)
		return
	}

	// Encode response
	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	playerID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(playerID)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	player, err := ph.playerRepo.GetPlayerByID(id)
	if err != nil {
		if err.Error() == "player not found" {
			http.Error(w, "Player not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(player)
}

func (ph *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		fmt.Println("Error decoding JSON request body:", err)
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Insert new player into the repository
	err = ph.playerRepo.CreatePlayer(&player)
	if err != nil {
		http.Error(w, "Failed to create player", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Player created successfully"))
}

func (ph *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(playerId)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	var updatedPlayer models.Player
	err = json.NewDecoder(r.Body).Decode(&updatedPlayer)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	updatedPlayer.ID = id
	err = ph.playerRepo.UpdatePlayer(&updatedPlayer)
	if err != nil {
		http.Error(w, "Failed to update player", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player updated successfully"))
}

func (ph *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	playerId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(playerId)

	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	err = ph.playerRepo.DeletePlayer(id)
	if err != nil {
		http.Error(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list_of_players/list_of_players.html", http.StatusSeeOther)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player deleted successfully"))

}
