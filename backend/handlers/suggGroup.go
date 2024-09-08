package handlers

import (
	"encoding/json"
	"net/http"
	"socialNetwork/pkg/services"
)

func SuggGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var groupId string
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&groupId)
		if err != nil {
			services.SendFront(w, "Invalid request body", 500)
			return
		}

		members, e := MemberService.GetMembership(groupId)
		if e != nil {
			services.SendFront(w, "error get member", 500)
			return
		}

		sugg, er := GroupeService.GetSuggGroup(members)
		if er != nil {
			services.SendFront(w, "Error get suggestion group", 500)
			return
		}

		services.SendFront(w, sugg, 200)
	}
}
