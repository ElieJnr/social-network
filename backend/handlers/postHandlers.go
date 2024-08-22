package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
	"strings"
)

// handler qui gère la récupération des posts avant de les encapsuler dans un objet JSON
func PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := PostService.GetAllPosts(w, r)
		if err != nil {
			fmt.Println("Error getting posts:", err)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "Failed to encode posts as JSON", http.StatusInternalServerError)
			return
		}
	}
}

// handler qui gère la création de post
func CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r)
	}
}

// se charge de récuperer les posts et de les inserer dans la base de données s'il sont correctes
func CreatePost(w http.ResponseWriter, r *http.Request) {
	postValue := CheckPost(w, r)

	if !postValue.Success {
		http.Error(w, postValue.Error, http.StatusBadRequest)
		return
	}

	err := PostService.InsertPost(postValue, w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// c'est une redirection temporaire
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}

// se charge de vérifier si les données du post sont correctes
func CheckPost(w http.ResponseWriter, r *http.Request) models.CheckResult {
	content := strings.TrimSpace(r.FormValue("thread"))
	privacy := r.FormValue("privacy")
	photoURL := utils.UploadImage(w, r, "post")

	var allowedUsers []string
	if privacy == "almost_private" {
		allowedUsers = r.Form["allowedUsers"] // doit contenir les Userid des utilisateurs autorisés à voir le post
	}

	isValid, validationError := utils.IsValidPost(content, privacy, photoURL, allowedUsers)
	if !isValid {
		return models.CheckResult{
			Success: false,
			Error:   validationError,
		}
	}

	return models.CheckResult{
		Success:      true,
		PhotoURL:     photoURL,
		Content:      content,
		Status:       privacy,
		AllowedUsers: allowedUsers,
	}
}
