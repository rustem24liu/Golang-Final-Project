package handlers

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"html/template"
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

func (ph *PlayerHandler) ListOfPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file
	tmpl, err := template.ParseFiles("cmd/list_of_players.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (ph *PlayerHandler) SortById(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.GetAllPlayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) SortByFirstname(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.SortByFirstname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) SortByLastname(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.SortByLastname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) SortByAge(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.SortByAge()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (ph *PlayerHandler) SortByCost(w http.ResponseWriter, r *http.Request) {
	players, err := ph.playerRepo.SortByCost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player deleted successfully"))

}
