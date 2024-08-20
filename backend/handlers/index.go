package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// un exemple pour verifier que notre api marche
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Cookie("session_token"))
		// Logique pour le handler de la route home
		json.NewEncoder(w).Encode([]byte("Welcome to the Social Network API!"))

	}
}
