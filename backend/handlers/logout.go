package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
)

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.CurrentUser(w, r)

		fmt.Println("salam", user)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = SessionService.DeleteSession(user.UserId)
		if err != nil {
			fmt.Println("err",err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		services.DelCookie(w)

	}
}
