package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
	"github.com/rustem24liu/Golang-Final-Project/internal/repository"
)

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("postgres", "postgres://postgres:1000tenge@localhost/football_team?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	playerRepo := repository.NewPlayerRepo(db)
	playerHandler := handlers.NewPlayerHandler(db)

	router.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		sortBy := r.URL.Query().Get("sort")
		positionFilter := r.URL.Query().Get("position") // Example filtering parameter

		filters := make(map[string]interface{})
		if positionFilter != "" {
			filters["player_pos"] = positionFilter
		}

		players, err := playerRepo.GetAllPlayers(pageNum, pageSize, sortBy, filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Encode response
		json.NewEncoder(w).Encode(players)
	},
	).Methods("GET")

	//router.HandleFunc('/')
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/tournament", handlers.TournamentHandler).Methods("GET")
	// Start HTTP server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
