package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/middlewares"
)

func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'utilisateur depuis le contexte
		user, ok := r.Context().Value(middlewares.UserContextKey).(*models.User)
		if !ok || user == nil {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		// Utiliser les informations de l'utilisateur
		fmt.Fprintf(w, user.Id.String(), user.Firstname, user.Email)
	}
}
