package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/services"

	"github.com/gofrs/uuid/v5"
)

type FollowRequest struct {
	UserId     string `json:"userId"`
	FollowedId string `json:"followedId"`
	Statut     bool   `json:"statut"`
}

func Follow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req FollowRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Println("one", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userId, err := uuid.FromString(req.UserId)
		if err != nil {
			fmt.Println("two", err)
			http.Error(w, "Cannot convert userId", http.StatusBadRequest)
			return
		}

		followedId, err := uuid.FromString(req.FollowedId)
		if err != nil {
			fmt.Println("tree", err)
			http.Error(w, "Cannot convert followedId", http.StatusBadRequest)
			return
		}

		followService := services.NewFollowerService()
		err = followService.FollowUserOrUpdateStatus(userId, followedId, req.Statut)
		if err != nil {
			fmt.Println("four", err)

			http.Error(w, "Cannot follow", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		services.SendFront(w, "sucess", http.StatusCreated)
	}
}
