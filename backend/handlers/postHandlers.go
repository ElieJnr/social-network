package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
	"strings"
)

// handler qui gère la récupération des posts avant de les encapsuler dans un objet JSON
func PostHandler(types string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if types == "userPost" {
			posts, err := PostService.GetOwnPosts(w, r)
			if err != nil {
				fmt.Println("Error getting user posts:", err)
				services.SendFront(w, map[string]string{"error": "Error getting user posts"}, http.StatusInternalServerError)
				return
			}
			if err := services.SendFront(w, posts, http.StatusOK); err != nil {
				fmt.Println("Failed to send user posts as JSON:", err)
			}
		} else {
			posts, err := PostService.GetAllPosts(w, r)
			if err != nil {
				fmt.Println("Error getting posts:", err)
				services.SendFront(w, map[string]string{"error": "Error getting posts"}, http.StatusInternalServerError)
				return
			}
			if err := services.SendFront(w, posts, http.StatusOK); err != nil {
				fmt.Println("Failed to send posts as JSON:", err)
			}
		}
	}
}

// handler qui gère la création de post
func CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r)
		// PostHandler("allPost")(w, r)
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
	
}

// se charge de vérifier si les données du post sont correctes
func CheckPost(w http.ResponseWriter, r *http.Request) models.CheckResult {
	content := strings.TrimSpace(r.FormValue("thread"))
	privacy := r.FormValue("privacy")
	photoURL, err := utils.UploadImage(w, r, "post")
	fmt.Println("err",err)

	if err != nil {
		return models.CheckResult{
			Success: false,
			Error:   err.Error(),
		}
	}
	// fmt.Println("check post values", content, privacy, photoURL)

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
