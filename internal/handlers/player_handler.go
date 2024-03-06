package handlers

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"github.com/rustem24liu/Golang-Final-Project/internal/repository"
	_ "github.com/rustem24liu/Golang-Final-Project/internal/repository"
	_ "log"
	"net/http"
	_ "net/http"
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
	// Extract player ID from request parameters
	// Retrieve player by ID from the repository
	// Encode player to JSON and write response
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
