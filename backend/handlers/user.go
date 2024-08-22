package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/utils"
)

func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'utilisateur depuis le contexte

		user, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}
		// Utiliser les informations de l'utilisateur
		fmt.Fprintf(w, user.Token, user.Username, user.UserId)
	}
}
