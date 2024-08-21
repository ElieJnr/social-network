package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/utils"
)

func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'utilisateur depuis le contexte

		user := utils.CurrentUser(w, r)
		// Utiliser les informations de l'utilisateur
		fmt.Fprintf(w, user.Token, user.Username, user.UserId)
	}
}
