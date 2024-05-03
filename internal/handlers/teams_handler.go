package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/rustem24liu/Golang-Final-Project/models"
	"net/http"

	"github.com/rustem24liu/Golang-Final-Project/internal/repository"
)

type TeamHandler struct {
	teamRepo *repository.TeamRepo
}

func NewTeamHandler(db *sql.DB) *TeamHandler {
	return &TeamHandler{
		teamRepo: repository.NewTeamRepo(db),
	}
}

func (ph *TeamHandler) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := ph.teamRepo.GetAllTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams)
}

// CreateTeamHandler creates a new team.
func (ph *TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ph.teamRepo.CreateTeam(r.Context(), &team); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}
