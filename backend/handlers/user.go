package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/middlewares"
	"socialNetwork/pkg/models"
)

func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'utilisateur depuis le contexte
		user, ok := r.Context().Value(middlewares.UserContextKey).(*models.Session)
		if !ok || user == nil {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		// Utiliser les informations de l'utilisateur
		fmt.Fprintf(w, user.Token, user.Username, user.UserId)
	}
}
