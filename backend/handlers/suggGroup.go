package handlers

import (
	"encoding/json"
	"net/http"
)

func SuggGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var groupId string
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&groupId)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)

			return
		}

		// members, e := MemberService.GetMembership(groupId)
		// http.Error(w, "Invalid request body", http.StatusBadRequest)
		// return
	}
}
