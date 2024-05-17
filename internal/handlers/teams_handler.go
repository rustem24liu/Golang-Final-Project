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

type TeamHandler struct {
	teamRepo *repository.TeamRepo
}

func NewTeamHandler(db *sql.DB) *TeamHandler {
	return &TeamHandler{
		teamRepo: repository.NewTeamRepo(db),
	}
}

func ListTeamsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database to fetch all players
	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold player data
	var teams []models.Team

	// Iterate over the rows and populate the players slice
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.TeamName, team.LeagueID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		teams = append(teams, team)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the list_of_players.html template with the player data
	tmpl, err := template.ParseFiles("/tmpl/list_of_teams.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, teams)
}

func (ph *TeamHandler) ListOfAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := ph.teamRepo.ListOfAllTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teams)
}

func (ph *TeamHandler) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	sortBy := r.URL.Query().Get("sort")
	nameFilter := r.URL.Query().Get("league_name")
	fmt.Println(nameFilter, "this is position filter")

	// Construct filter map
	filters := make(map[string]interface{})
	if nameFilter != "" {
		filters["league_name"] = nameFilter
	}

	// Retrieve players from repository
	teams, err := ph.teamRepo.GetAllTeams(pageNum, pageSize, sortBy, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error retrieving team:", err)
		return
	}

	// Encode response
	json.NewEncoder(w).Encode(teams)
}

func (ph *TeamHandler) GetTeamByID(w http.ResponseWriter, r *http.Request) {
	teamID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	team, err := ph.teamRepo.GetTeamByID(id)
	if err != nil {
		if err.Error() == "team not found" {
			http.Error(w, "team not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(team)
}

func (ph *TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		fmt.Println("Error decoding JSON request body:", err)
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Insert new player into the repository
	err = ph.teamRepo.CreateTeam(&team)
	if err != nil {
		http.Error(w, "Failed to create team", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Team created successfully"))
}

func (ph *TeamHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	teamId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(teamId)
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	var updatedTeam models.Team
	err = json.NewDecoder(r.Body).Decode(&updatedTeam)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	updatedTeam.ID = id
	err = ph.teamRepo.UpdateTeam(&updatedTeam)
	if err != nil {
		http.Error(w, "Failed to update team", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("League updated successfully"))
}

func (ph *TeamHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	teamId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(teamId)

	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	err = ph.teamRepo.DeleteTeam(id)
	if err != nil {
		http.Error(w, "Failed to delete team", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tmpl/list_of_team.html", http.StatusSeeOther)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player deleted successfully"))

}
