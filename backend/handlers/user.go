package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"

	"github.com/gofrs/uuid/v5"
)

type Info struct {
	UserId  string `json:"userId"`
	Private bool   `json:"private"`
}

func UsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userService := services.NewUserService()

		users, err := userService.GetAllUsers(w,r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Cannot get users ", http.StatusInternalServerError)
			return
		}
		services.SendFront(w, map[string][]models.User{"users": users}, 200)
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
		if key != "userConnect" {
			id, err = uuid.FromString(query.Get("id"))
			if err != nil {
				http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
				return
			}
		}
		userService := services.NewUserService()
		user, err := userService.GetUserById(w,r,id)
		if err != nil {
			fmt.Println("Error", err)
			http.Error(w, "Cannot convert to uuid", http.StatusBadRequest)
			return
		}
		services.SendFront(w, map[string]models.User{"user": user}, 200)
	}
}

func Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var req Info
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		userId, err := uuid.FromString(req.UserId)
		if err != nil {
			http.Error(w, "Cannot convert userId", http.StatusBadRequest)
			return
		}
		userService := services.NewUserService()

		err = userService.UpdateUser(req.Private, userId)
		if err != nil {
			http.Error(w, "Cannot update userId", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		services.SendFront(w, "sucess", http.StatusCreated)
	}
}
