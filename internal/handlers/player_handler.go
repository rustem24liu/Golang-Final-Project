package handlers

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rustem24liu/Golang-Final-Project/internal/repository"
	_ "github.com/rustem24liu/Golang-Final-Project/internal/repository"
	_ "log"
	"net/http"
	_ "net/http"
	"strconv"
)

type PlayerHandler struct {
	playerRepo *repository.PlayerRepo
}

func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	return &PlayerHandler{
		playerRepo: repository.NewPlayerRepo(db),
	}
}

func (ph *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.GetAllPlayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	playerID := mux.Vars(r)["id"]

	fmt.Println(playerID)

	id, err := strconv.Atoi(playerID)
	if err != nil {
		fmt.Println(playerID, "smth bad")
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	player, err := ph.playerRepo.GetPlayerById(id)
	if err != nil {
		if err.Error() == "player not found" {
			http.Error(w, "Player not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return player details in JSON format
	json.NewEncoder(w).Encode(player)
}

func (ph *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body to Player struct
	// Insert new player into the repository
	// Write success response or error if any occurs
}

func (ph *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	// Extract player ID from request parameters
	// Decode JSON request body to Player struct
	// Update player in the repository
	// Write success response or error if any occurs
}

func (ph *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// Extract player ID from request parameters
	// Delete player from the repository
	// Write success response or error if any occurs
}
