package handlers

import (
	"encoding/json"
	"net/http"
	"socialNetwork/pkg/models"
)

func AddNewMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Vérifie que la méthode est bien POST
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		// Décode les données JSON envoyées par le frontend
		var newMember models.NewMember

		err := json.NewDecoder(r.Body).Decode(&newMember)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Insère les données dans la base de données

		newMember.UserId = "56cec81e-e5b3-432d-8f3b-90d0cc29b09d"
		newMember.Status="member"

		if newMember.GroupId == "" || newMember.UserId == "" || newMember.Status=="" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err = GroupeService.AddNewMember(newMember)

		if err != nil {
			http.Error(w, "failed to add new member in the group", http.StatusUnauthorized)
			return
		}

		// Réponse avec succès
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("New member added successfully"))
	}
}
