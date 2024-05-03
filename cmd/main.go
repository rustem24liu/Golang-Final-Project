package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
)

const (
	portEnvKey   = "PORT"
	hostEnvKey   = "DB_HOST"
	portEnvName  = 5432
	userEnvKey   = "DB_USER"
	passEnvKey   = "DB_PASSWORD"
	dbnameEnvKey = "DB_NAME"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the 404 HTML file
	http.ServeFile(w, r, "cmd/404/404.html")
}

func main() {

	port := os.Getenv(portEnvKey)
	if port == "" {
		port = "8080" // Default port if not specified
	}

	router := mux.NewRouter()
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Database Connection String:", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	playerHandler := handlers.NewPlayerHandler(db)
	//teamRepo := repository.NewTeamRepo(db)
	teamHandler := handlers.NewTeamHandler(db)

	handlers.SetDB(db)
	//router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods("GET")
	router.HandleFunc("/teams", teamHandler.CreateTeam).Methods("POST")
	router.HandleFunc("/activate", handlers.ActivateUserHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.Handle("/protected", handlers.Authenticate(http.HandlerFunc(handlers.ProtectedHandler))).Methods("GET")
	//router.HandleFunc('/')
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/players", playerHandler.GetAllPlayers).Methods("GET")
	router.HandleFunc("/players/list", playerHandler.ListOfAllPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/tournament", handlers.TournamentHandler).Methods("GET")
	//router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods("GET")
	router.HandleFunc("/developers", handlers.DevelopersHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.HandleFunc("/players/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListPlayersHandler(w, r, db)
	})

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	pageNum, _ := strconv.Atoi(queryParams.Get("page"))

	pageSize, _ := strconv.Atoi(queryParams.Get("size"))

	sortBy := queryParams.Get("sort")

	positionFilter := queryParams.Get("player_pos")
	fmt.Fprintf(w, "pageNum: %d, pageSize: %d, sortBy: %s, positionFilter: %s", pageNum, pageSize, sortBy, positionFilter)
}
