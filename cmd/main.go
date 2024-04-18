package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/handlers"
)

var jwtKey = []byte("secret_key")

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var users = []User{
	{ID: 1, Username: "user1", Password: "password1"},
	{ID: 2, Username: "user2", Password: "password2"},
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var foundUser User
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			foundUser = user
			break
		}
	}

	if foundUser.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &JWTClaims{
		Username: foundUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", tokenString)
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the protected area!"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 page not found")
	})

	// Protected endpoint
	router.Handle("/protected", authenticate(http.HandlerFunc(protectedHandler))).Methods("GET")
	db, err := sql.Open("postgres", "postgres://postgres:1000tenge@localhost/football_team?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	//router.HandleFunc('/')

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/players", playerHandler.GetAllPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerHandler.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/tournament", handlers.TournamentHandler).Methods("GET")
	router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods("GET")
	// Start HTTP server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
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
