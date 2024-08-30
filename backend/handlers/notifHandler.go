package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
)

func NotifHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.CurrentUser(w, r)

		if err != nil {
			fmt.Println("1")
			services.SendFront(w, models.Errors["500"], 500)
			return
		}
		fmt.Println(user.UserId)
		notis, err := NotifService.GetNotifications(user.UserId)
		fmt.Println(notis)
		if err != nil {
			fmt.Println("2", err)
			services.SendFront(w, models.Errors["500"], 500)
			return
		}

		services.SendFront(w, notis, 200)
	}
}
