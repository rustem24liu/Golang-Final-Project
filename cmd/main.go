package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ayan2004"
	dbname   = "football_teams"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the 404 HTML file
	http.ServeFile(w, r, "404/404.html") // Update with the actual path to your 404 HTML file
}

func main() {

	router := mux.NewRouter()
	//router.HandleFunc("/login", loginHandler).Methods("POST")
	//router.HandleFunc("/register", registerHandler).Methods("POST")
	//
	//router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusNotFound)
	//	fmt.Fprintf(w, "404 page not found")
	//})
	//
	//// Protected endpoint
	//router.Handle("/protected", authenticate(http.HandlerFunc(protectedHandler))).Methods("GET")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Connect to the PostgreSQL database

	playerHandler := handlers.NewPlayerHandler(db)
	teamHandler := handlers.NewTeamHandler(db)

	//playerRepo := repository.NewPlayerRepo(db)
	//playerHandler := handlers.NewPlayerHandler(db)

	//router.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
	//	pageNum, _ := strconv.Atoi(r.URL.Query().Get("page"))
	//	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	//	sortBy := r.URL.Query().Get("sort")
	//	positionFilter := r.URL.Query().Get("position") // Example filtering parameter
	//
	//	filters := make(map[string]interface{})
	//	if positionFilter != "" {
	//		filters["player_pos"] = positionFilter
	//	}
	//
	//	players, err := playerRepo.GetAllPlayers(pageNum, pageSize, sortBy, filters)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		fmt.Println("some error main")
	//		return
	//	}
	//	// Encode response
	//	json.NewEncoder(w).Encode(players)
	//},
	//).Methods("GET")
	handlers.SetDB(db)
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
	router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods("GET")
	router.HandleFunc("/developers", handlers.DevelopersHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.HandleFunc("/players/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListPlayersHandler(w, r, db)
	})

	// Start HTTP server
	port := 8080
	fmt.Printf("http://localhost:8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL
	queryParams := r.URL.Query()

	// Parse pageNum
	pageNum, _ := strconv.Atoi(queryParams.Get("page"))

	// Parse pageSize
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))

	// Parse sortBy
	sortBy := queryParams.Get("sort")

	// Parse filters
	positionFilter := queryParams.Get("player_pos")

	// Now you have pageNum, pageSize, sortBy, and positionFilter
	// You can pass these values to your handler function or use them as needed
	// For example, you can call your handler function with these parameters:
	// YourHandlerFunction(w, r, pageNum, pageSize, sortBy, positionFilter)
	fmt.Fprintf(w, "pageNum: %d, pageSize: %d, sortBy: %s, positionFilter: %s", pageNum, pageSize, sortBy, positionFilter)
}
