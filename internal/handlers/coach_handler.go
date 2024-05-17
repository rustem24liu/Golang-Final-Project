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

type CoachHandler struct {
	coachRepo *repository.CoachRepo
}

func NewCoachHandler(db *sql.DB) *CoachHandler {
	return &CoachHandler{
		coachRepo: repository.NewCoachRepo(db),
	}
}

func (ph *CoachHandler) ListOfAllCoaches(w http.ResponseWriter, r *http.Request) {
	coaches, err := ph.coachRepo.ListOfAllCoaches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coaches)
}

func (ph *CoachHandler) GetAllCoaches(w http.ResponseWriter, r *http.Request) {
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	sortBy := r.URL.Query().Get("sort")
	positionFilter := r.URL.Query().Get("first_name")
	fmt.Println(positionFilter, "this is position filter")

	// Construct filter map
	filters := make(map[string]interface{})
	if positionFilter != "" {
		filters["first_name"] = positionFilter
	}

	// Retrieve players from repository
	coaches, err := ph.coachRepo.GetAllCoaches(pageNum, pageSize, sortBy, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error retrieving coach:", err)
		return
	}

	// Encode response
	json.NewEncoder(w).Encode(coaches)
}

func (ph *CoachHandler) GetCoachByID(w http.ResponseWriter, r *http.Request) {
	coachID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(coachID)
	if err != nil {
		http.Error(w, "Invalid coach ID", http.StatusBadRequest)
		return
	}

	coach, err := ph.coachRepo.GetCoachByID(id)
	if err != nil {
		if err.Error() == "coach not found" {
			http.Error(w, "coach not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coach)
}

func (ph *CoachHandler) CreateCoach(w http.ResponseWriter, r *http.Request) {
	var coach models.Coach

	err := json.NewDecoder(r.Body).Decode(&coach)
	if err != nil {
		fmt.Println("Error decoding JSON request body:", err)
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Insert new player into the repository
	err = ph.coachRepo.CreateCoach(&coach)
	if err != nil {
		http.Error(w, "Failed to create coach", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Coach created successfully"))
}

func (ph *CoachHandler) UpdateCoach(w http.ResponseWriter, r *http.Request) {
	coachId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(coachId)
	if err != nil {
		http.Error(w, "Invalid coach ID", http.StatusBadRequest)
		return
	}

	var updatedCoach models.Coach
	err = json.NewDecoder(r.Body).Decode(&updatedCoach)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	updatedCoach.ID = id
	err = ph.coachRepo.UpdateCoach(&updatedCoach)
	if err != nil {
		http.Error(w, "Failed to update coach", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Coach updated successfully"))
}

func (ph *CoachHandler) DeleteCoach(w http.ResponseWriter, r *http.Request) {
	coachId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(coachId)

	if err != nil {
		http.Error(w, "Invalid coach ID", http.StatusBadRequest)
		return
	}

	err = ph.coachRepo.DeleteCoach(id)
	if err != nil {
		http.Error(w, "Failed to delete coach", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "tmpl/list_of_coaches.html", http.StatusSeeOther)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Coach deleted successfully"))

}
