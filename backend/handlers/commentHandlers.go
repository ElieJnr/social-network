package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
	"strings"
)

func CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateComment(w, r)
	}
}

func CommentHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postId := r.URL.Query().Get("postId") 

        comments, err := CommentService.GetComments(postId)
        if err != nil {
            fmt.Println("Error getting comments:", err)
            services.SendFront(w, map[string]string{"error": "Error getting comments"}, http.StatusInternalServerError)
            return
        }
        if err := services.SendFront(w, comments, http.StatusOK); err != nil {
            fmt.Println("Failed to send comments as JSON:", err)
        }
    }
}


func CreateComment(w http.ResponseWriter, r *http.Request) {
	commentValue := CheckComment(w, r)

	if !commentValue.Success {
		fmt.Println("Error: ", commentValue.Error)
		http.Error(w, commentValue.Error, http.StatusBadRequest)
		return
	}

	err := CommentService.InsertComment(commentValue, w, r)
	if err != nil {
		fmt.Println("Error InsertComment: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func CheckComment(w http.ResponseWriter, r *http.Request) models.CheckResult {
	content := strings.TrimSpace(r.FormValue("commentContent"))
	postId := r.FormValue("postId")
	photoURL, errIMage := utils.UploadImage(w, r, "comment")
	if errIMage != nil {
		return models.CheckResult{
			Success: false,
			Error:   errIMage.Error(),
		}
	}

	success, err := utils.IsValidComment(content, photoURL, postId)
	if !success {
		return models.CheckResult{
			Success: false,
			Error:   err,
		}
	}

	return models.CheckResult{
		Success:  true,
		Error:    "",
		PostId:   postId,
		Content:  content,
		PhotoURL: photoURL,
	}
}
