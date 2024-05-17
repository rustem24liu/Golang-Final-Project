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

type StadiumHandler struct {
	stadiumRepo *repository.StadiumRepo
}

func NewStadiumHandler(db *sql.DB) *StadiumHandler {
	return &StadiumHandler{
		stadiumRepo: repository.NewStadiumRepo(db),
	}
}

func ListStadiumsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Query the database to fetch all players
	rows, err := db.Query("SELECT * FROM Stadium")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold player data
	var stadiums []models.Stadium

	// Iterate over the rows and populate the players slice
	for rows.Next() {
		var stadium models.Stadium
		if err := rows.Scan(&stadium.StadiumName, &stadium.Capacity, &stadium.TeamID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stadiums = append(stadiums, stadium)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the list_of_players.html template with the player data
	tmpl, err := template.ParseFiles("tmpl/list_of_stadiums.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, stadiums)
}

func (ph *StadiumHandler) ListOfAllStadiums(w http.ResponseWriter, r *http.Request) {
	stadiums, err := ph.stadiumRepo.ListOfAllStadiums()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stadiums)
}

func (ph *StadiumHandler) GetAllStadiums(w http.ResponseWriter, r *http.Request) {
	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	sortBy := r.URL.Query().Get("sort")
	positionFilter := r.URL.Query().Get("stadium_name")
	fmt.Println(positionFilter, "this is position filter")

	// Construct filter map
	filters := make(map[string]interface{})
	if positionFilter != "" {
		filters["stadium_name"] = positionFilter
	}

	// Retrieve players from repository
	stadiums, err := ph.stadiumRepo.GetAllStadiums(pageNum, pageSize, sortBy, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error retrieving stadium:", err)
		return
	}

	// Encode response
	json.NewEncoder(w).Encode(stadiums)
}

func (ph *StadiumHandler) GetStadiumByID(w http.ResponseWriter, r *http.Request) {
	stadiumID := mux.Vars(r)["id"]

	id, err := strconv.Atoi(stadiumID)
	if err != nil {
		http.Error(w, "Invalid stadium ID", http.StatusBadRequest)
		return
	}

	stadium, err := ph.stadiumRepo.GetStadiumByID(id)
	if err != nil {
		if err.Error() == "stadium not found" {
			http.Error(w, "stadium not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stadium)
}

func (ph *StadiumHandler) CreateStadium(w http.ResponseWriter, r *http.Request) {
	var stadium models.Stadium

	err := json.NewDecoder(r.Body).Decode(&stadium)
	if err != nil {
		fmt.Println("Error decoding JSON request body:", err)
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Insert new player into the repository
	err = ph.stadiumRepo.CreateStadium(&stadium)
	if err != nil {
		http.Error(w, "Failed to create stadium", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Stadium created successfully"))
}

func (ph *StadiumHandler) UpdateStadium(w http.ResponseWriter, r *http.Request) {
	stadiumId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(stadiumId)
	if err != nil {
		http.Error(w, "Invalid stadium ID", http.StatusBadRequest)
		return
	}

	var updatedStadium models.Stadium
	err = json.NewDecoder(r.Body).Decode(&updatedStadium)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	updatedStadium.ID = id
	err = ph.stadiumRepo.UpdateStadium(&updatedStadium)
	if err != nil {
		http.Error(w, "Failed to update stadium", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Coach updated successfully"))
}

func (ph *StadiumHandler) DeleteStadium(w http.ResponseWriter, r *http.Request) {
	stadiumId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(stadiumId)

	if err != nil {
		http.Error(w, "Invalid stadium ID", http.StatusBadRequest)
		return
	}

	err = ph.stadiumRepo.DeleteStadium(id)
	if err != nil {
		http.Error(w, "Failed to delete stadium", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "tmpl/list_of_stadiums.html", http.StatusSeeOther)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stadium deleted successfully"))

}
