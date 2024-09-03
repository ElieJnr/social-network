package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var userID = "56cec81e-e5b3-432d-8f3b-90d0cc29b09d"
		groups, err := GroupeService.GetGroups(userID)

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
