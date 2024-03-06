package main

import (
	"database/sql"
	_ "fmt"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
	"log"
	"net/http"
)

func main() {
	//tournament.RunTournament()
	//fmt.Println("Tournament finished")
	//
	//rows, err := database.GetTeamNames()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rows)

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/football_team?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize player handler
	playerHandler := handlers.NewPlayerHandler(db)

	http.HandleFunc("/players", playerHandler.GetAllPlayers)
	http.HandleFunc("/players/:id", playerHandler.GetPlayerByID)
	http.HandleFunc("/players/create", playerHandler.CreatePlayer)
	http.HandleFunc("/players/update/:id", playerHandler.UpdatePlayer)
	http.HandleFunc("/players/delete/:id", playerHandler.DeletePlayer)

	// Start HTTP server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
