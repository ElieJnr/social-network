package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"

	"github.com/gofrs/uuid/v5"
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

func GetUserConnectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'utilisateur depuis le contexte
		userInfo, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}
		fmt.Println(userInfo)
		id, err := uuid.FromString(userInfo.UserId)
		if err != nil {
			http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
			return
		}
		userService := services.NewUserService()
		user, err := userService.GetUserById(id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
			return
		}
		services.SendFront(w, map[string]models.User{"user": user}, 200)
	}
}
