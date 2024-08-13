package handlers

import (
	"net/http"
	"socialNetwork/pkg/db/sqlite"
)

// un exemple pour verifier que notre api marche
func HomeHandler(db *sqlite.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logique pour le handler de la route home
		w.Write([]byte("Welcome to the Social Network API!"))
	}
}
