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
	teamHandler := handlers.NewTeamHandler(db)
	stadiumHandler := handlers.NewStadiumHandler(db)
	coachHandler := handlers.NewCoachHandler(db)

	leagueHandler := handlers.NewLeagueHandler(db)

	handlers.SetDB(db)

	// authorization & authentication
	router.HandleFunc("/activate", handlers.ActivateUserHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.Handle("/protected", handlers.Authenticate(http.HandlerFunc(handlers.ProtectedHandler))).Methods("GET")

	// main page
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// additional logics
	router.HandleFunc("/tournament", handlers.TournamentHandler).Methods("GET")
	router.HandleFunc("/developers", handlers.DevelopersHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// players - CRUD
	router.HandleFunc("/players", playerHandler.GetAllPlayers).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/list", playerHandler.ListOfAllPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")

	// team - CRUD
	router.HandleFunc("/teams", teamHandler.CreateTeam).Methods("POST")
	router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", teamHandler.GetTeamByID).Methods("GET")
	router.HandleFunc("/teams/{id}", teamHandler.UpdateTeam).Methods("PUT")
	router.HandleFunc("/teams/{id}", teamHandler.DeleteTeam).Methods("DELETE")

	// stadium - CRUD
	router.HandleFunc("/stadiums", stadiumHandler.CreateStadium).Methods("POST")
	router.HandleFunc("/stadiums", stadiumHandler.GetAllStadiums).Methods("GET")
	router.HandleFunc("/stadiums/{id}", stadiumHandler.GetStadiumByID).Methods("GET")
	router.HandleFunc("/stadiums/{id}", stadiumHandler.UpdateStadium).Methods("PUT")
	router.HandleFunc("/stadiums/{id}", stadiumHandler.DeleteStadium).Methods("DELETE")

	// coach - CRUD
	router.HandleFunc("/coaches", coachHandler.CreateCoach).Methods("POST")
	router.HandleFunc("/coaches", coachHandler.GetAllCoaches).Methods("GET")
	router.HandleFunc("/coaches/{id}", coachHandler.GetCoachByID).Methods("GET")
	router.HandleFunc("/coaches/{id}", coachHandler.UpdateCoach).Methods("PUT")
	router.HandleFunc("/coaches/{id}", coachHandler.DeleteCoach).Methods("DELETE")

	// league - CRUD
	router.HandleFunc("/league", leagueHandler.CreateLeague).Methods("POST")
	router.HandleFunc("/league", leagueHandler.GetAllLeagues).Methods("GET")
	router.HandleFunc("/league/{id}", leagueHandler.GetLeagueByID).Methods("GET")
	router.HandleFunc("/league/{id}", leagueHandler.UpdateLeague).Methods("PUT")
	router.HandleFunc("/league/{id}", leagueHandler.DeleteLeague).Methods("DELETE")

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
