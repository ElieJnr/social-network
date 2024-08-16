package handlers

import (
	"net/http"
)

// un exemple pour verifier que notre api marche
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logique pour le handler de la route home
		w.Write([]byte("Welcome to the Social Network API!"))
	}
}
