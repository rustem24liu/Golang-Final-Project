package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
	"github.com/rustem24liu/Golang-Final-Project/internal/tournament"
)

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("postgres", "postgres://postgres:1000tenge@localhost/football_team?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize player handler
	playerHandler := handlers.NewPlayerHandler(db)
	tournament.RunTournament()

	router.HandleFunc("/players", playerHandler.GetAllPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")

	// Start HTTP server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
