package handlers

import (
	"net/http"
)

// DevelopersHandler serves the developers.html file
// DevelopersHandler serves the developers.html file
func DevelopersHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file
	http.ServeFile(w, r, "cmd/developers/developers.html") // Update with the actual path to developers.html
}
