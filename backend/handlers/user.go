package handlers

import (
	"encoding/json"
	"net/http"
)

// un exemple pour retouner du json
func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logique pour gérer les utilisateurs
		users := []string{"user1", "user2", "user3"}
		json.NewEncoder(w).Encode(users)
	}
}
