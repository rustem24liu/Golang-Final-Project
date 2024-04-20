package handlers

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"strings"

	"github.com/dgrijalva/jwt-go"
)
var db *sql.DB

// SetDB sets the database connection
func SetDB(database *sql.DB) {
    db = database
}

var jwtKey = []byte("secret_key")

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Activated    bool   `json:"activated"`
	Permissions  []string `json:"permissions"`
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
	{ID: 1, Username: "user1", Password: "password1", Activated: true, Permissions: []string{"read"}},
	{ID: 2, Username: "user2", Password: "password2", Activated: false, Permissions: []string{"read"}},
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Check if the username is already taken
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", newUser.Username).Scan(&count)
    if err != nil {
        http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
        return
    }
    if count > 0 {
        http.Error(w, "Username is already taken", http.StatusBadRequest)
        return
    }

    // Convert permissions slice to a formatted string for PostgreSQL array literal
    permissionsStr := "{" + strings.Join(newUser.Permissions, ",") + "}"

    // Insert the new user into the database
    _, err = db.Exec("INSERT INTO users (username, password, activated, permissions) VALUES ($1, $2, $3, $4)",
        newUser.Username, newUser.Password, newUser.Activated, permissionsStr)
    if err != nil {
        http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var foundUser *User
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			foundUser = &user
			break
		}
	}

	if foundUser == nil || !foundUser.Activated {
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

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims := &JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Add permission-based authorization logic here
		if !checkPermissions(claims.Username, r.URL.Path) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkPermissions(username string, path string) bool {
	// Placeholder logic for checking user permissions based on username and path
	for _, user := range users {
		if user.Username == username {
			for _, permission := range user.Permissions {
				if permission == "admin" && path == "/admin" {
					return true
				} else if permission == "read" && path == "/data" {
					return true
				}
			}
		}
	}
	return false
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the protected area!"))
}
