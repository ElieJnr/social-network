package handlers

import (
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
	"strings"
)

func CreateCommentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateComment(w, r)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var postId string
	commentValue := CheckComment(w, r)

	if !commentValue.Success {
		http.Error(w, commentValue.Error, http.StatusBadRequest)
		return
	}

	err := CommentService.InsertComment(commentValue, w, r, postId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}


func CheckComment(w http.ResponseWriter, r *http.Request) models.CheckResult {
	content := strings.TrimSpace(r.FormValue("commentContent"))
	photoURL := utils.UploadImage(w, r, "comment")

	success, err := utils.IsValidComment(content, photoURL)
	if !success {
		return models.CheckResult{
			Success: false,
			Error:   err,
		}
	}

	return models.CheckResult{
		Success:  true,
		Error:    "",
		Content:  content,
		PhotoURL: photoURL,
	}
}
