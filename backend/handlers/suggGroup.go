package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/services"
)

func SuggGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ooooooooooooooooooooo")
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var groupId string
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&groupId)
		fmt.Println(groupId)
		if err != nil {
			fmt.Println("1")
			services.SendFront(w, "Invalid request body", 500)
			return
		}

		members, e := MemberService.GetMembership(groupId)
		if e != nil {
			fmt.Println("2")
			services.SendFront(w, "error get member", 500)
			return
		}

		sugg, er := GroupeService.GetSuggGroup(members)
		if er != nil {
			fmt.Println(er)
			services.SendFront(w, "Error get suggestion group", 500)
			return
		}
		fmt.Println(sugg)
		services.SendFront(w, sugg, 200)
	}
}
