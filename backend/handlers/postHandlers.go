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
			userId := r.URL.Query().Get("userId")
			posts, err := PostService.GetPostsById(w, r, userId)
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
		CreatePost(w, r, "post")
	}
}

// handler qui gère la création de post pour un groupe
func GroupCreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r, "group")
	}
}

// handler qui gère la récupération des posts d'un groupe
func GroupPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId := r.URL.Query().Get("groupId")
		posts, err := PostService.GetGroupPosts(w, r, groupId)
		if err != nil {
			fmt.Println("Error getting group posts:", err)
			services.SendFront(w, map[string]string{"error": "Error getting group posts"}, http.StatusInternalServerError)
			return
		}
		if err := services.SendFront(w, posts, http.StatusOK); err != nil {
			fmt.Println("Failed to send group posts as JSON:", err)
		}
	}
}

// se charge de récuperer les posts et de les inserer dans la base de données s'il sont correctes
func CreatePost(w http.ResponseWriter, r *http.Request, origin string) {
	postValue := CheckPost(w, r, origin)
	if !postValue.Success {
		http.Error(w, postValue.Error, http.StatusBadRequest)
		return
	}

	err := PostService.InsertPost(postValue, w, r, origin)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// se charge de vérifier si les données du post sont correctes
func CheckPost(w http.ResponseWriter, r *http.Request, origin string) models.CheckResult {
	r.ParseForm()
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return models.CheckResult{
			Success: false,
			Error:   "Failed to parse form data: " + err.Error(),
		}
	}
	content := strings.TrimSpace(r.FormValue("thread"))
	photoURL, err := utils.UploadImage(w, r, origin)
	if err != nil {
		fmt.Println("Error uploading image:", err)
		return models.CheckResult{
			Success: false,
			Error:   err.Error(),
		}
	}

	var privacy string
	var allowedUsers []string
	var groupId string

	if origin == "post" {
		privacy = r.FormValue("privacy")
		if privacy == "almost-private" {
			allowedUsersStr := r.FormValue("allowedUsers")
			allowedUsers = strings.Split(allowedUsersStr, ",")
		}
	}

	if origin == "group" {
		groupId = r.FormValue("groupId")
	}

	// Validation du contenu, de la confidentialité et des utilisateurs autorisés
	isValid, validationError := utils.IsValidPost(content, privacy, photoURL, allowedUsers, origin)
	if !isValid {
		fmt.Println("Invalid post:", validationError)
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
		GroupId:      groupId,
		AllowedUsers: allowedUsers,
	}
}
