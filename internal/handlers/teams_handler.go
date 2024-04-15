package handlers

import (
	"database/sql"
	"encoding/json"
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
