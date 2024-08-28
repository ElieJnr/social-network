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

func GetUsertHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		userInfo, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}
		id, err := uuid.FromString(userInfo.UserId)
		if err != nil {
			http.Error(w, "Cannot convert ", http.StatusBadRequest)
			return
		}
		key := query.Get("key")
		if (key != "userConnect"){
			id, err = uuid.FromString(query.Get("id"))
			if err != nil {
				http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
				return
			} 
		}
		userService := services.NewUserService()
		user, err := userService.GetUserById(id)
		if err != nil {
			fmt.Println("Error",err)
			http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
			return
		}
		services.SendFront(w, map[string]models.User{"user": user}, 200)
	}
}
