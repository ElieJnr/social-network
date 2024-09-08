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
	FollowU    bool   `json:"followU"`
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
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userId, err := uuid.FromString(req.UserId)
		if err != nil {
			http.Error(w, "Cannot convert userId", http.StatusBadRequest)
			return
		}

		followedId, err := uuid.FromString(req.FollowedId)
		if err != nil {
			http.Error(w, "Cannot convert followedId", http.StatusBadRequest)
			return
		}

		followService := services.NewFollowerService()
		fmt.Println("requet----------------",req.FollowU)
		if req.FollowU {
			err = followService.FollowUserOrUpdateStatus(userId, followedId, req.Statut)
			if err != nil {
				http.Error(w, "Cannot follow", http.StatusInternalServerError)
				return
			}
		} else {
		
			err = followService.UnfollowUser(userId, followedId, req.Statut)
			if err != nil {
				http.Error(w, "Cannot follow", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		services.SendFront(w, "sucess", http.StatusCreated)
	}
}
