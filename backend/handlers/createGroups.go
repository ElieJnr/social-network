package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

// Fonction pour créer des groupes
func CreateGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Vérifie que la méthode est bien POST
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		// Décode les données JSON envoyées par le frontend
		var group models.Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Insère les données dans la base de données

		user, err := utils.CurrentUser(w, r)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		group.UserId = user.UserId

		if group.Title == "" || group.Description == "" || group.UserId == "" {
			fmt.Println("mboldé")
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		group.Id, err = GroupeService.CreateGroup(group)

		if err != nil {
			http.Error(w, "Creation of group failed", http.StatusUnauthorized)
			return
		}

		newMember := models.NewMember{
			UserId:  group.UserId,
			GroupId: group.Id,
			Status:  "admin",
		}

		err = GroupeService.AddNewMember(newMember)

		if err != nil {
			http.Error(w, "failed to add new member in the group", http.StatusUnauthorized)
			return
		}

		// Réponse avec succès
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Group created successfully"))
	}
}
