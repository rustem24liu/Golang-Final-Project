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

type LeagueHandler struct {
	leagueRepo *repository.LeagueRepo
}

func NewLeagueHandler(db *sql.DB) *LeagueHandler {
	return &LeagueHandler{
		leagueRepo: repository.NewLeagueRepo(db),
	}
}

func (ph *LeagueHandler) ListOfAllLeagues(w http.ResponseWriter, r *http.Request) {
	leagues, err := ph.leagueRepo.ListOfAllLeagues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(leagues)
}

func (ph *LeagueHandler) GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	sortBy := r.URL.Query().Get("sort")
	positionFilter := r.URL.Query().Get("league_name")
	fmt.Println(positionFilter, "this is position filter")

	// Construct filter map
	filters := make(map[string]interface{})
	if positionFilter != "" {
		filters["league_name"] = positionFilter
	}

	// Retrieve players from repository
	leagues, err := ph.leagueRepo.GetAllLeagues(pageNum, pageSize, sortBy, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error retrieving league:", err)
		return
	}

	// Encode response
	json.NewEncoder(w).Encode(leagues)
}

func (ph *LeagueHandler) GetLeagueByID(w http.ResponseWriter, r *http.Request) {
	leagueID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(leagueID)
	if err != nil {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	league, err := ph.leagueRepo.GetLeagueByID(id)
	if err != nil {
		if err.Error() == "league not found" {
			http.Error(w, "league not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(league)
}

func (ph *LeagueHandler) CreateLeague(w http.ResponseWriter, r *http.Request) {
	var league models.League

	err := json.NewDecoder(r.Body).Decode(&league)
	if err != nil {
		fmt.Println("Error decoding JSON request body:", err)
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Insert new player into the repository
	err = ph.leagueRepo.CreateLeague(&league)
	if err != nil {
		http.Error(w, "Failed to create league", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("League created successfully"))
}

func (ph *LeagueHandler) UpdateLeague(w http.ResponseWriter, r *http.Request) {
	leagueId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(leagueId)
	if err != nil {
		http.Error(w, "Invalid league ID", http.StatusBadRequest)
		return
	}

	var updatedLeague models.League
	err = json.NewDecoder(r.Body).Decode(&updatedLeague)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	updatedLeague.ID = id
	err = ph.leagueRepo.UpdateLeague(&updatedLeague)
	if err != nil {
		http.Error(w, "Failed to update league", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("League updated successfully"))
}

func (ph *LeagueHandler) DeleteLeague(w http.ResponseWriter, r *http.Request) {
	leagueId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(leagueId)

	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	err = ph.leagueRepo.DeleteLeague(id)
	if err != nil {
		http.Error(w, "Failed to delete league", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tmpl/list_of_players.html", http.StatusSeeOther)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("League deleted successfully"))

}
