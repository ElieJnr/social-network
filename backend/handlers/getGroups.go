package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/utils"
)

func GetGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userInfo, err := utils.CurrentUser(w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		groups, err := GroupeService.GetGroups(userInfo.UserId)		
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "impossible to get all the group", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(groups); err != nil {
			fmt.Println("Error encoding groups to JSON:", err)
			http.Error(w, "Failed to encode groups to JSON", http.StatusInternalServerError)
			return
		}
	}
}
