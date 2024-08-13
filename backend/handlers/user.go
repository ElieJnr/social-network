package handlers

import (
	"encoding/json"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
)

// un exemple pour retouner du json
func UsersHandler(db *sqlite.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logique pour gérer les utilisateurs
		users := []string{"user1", "user2", "user3"}
		json.NewEncoder(w).Encode(users)
	}
}
