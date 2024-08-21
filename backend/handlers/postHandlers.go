package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
	"strings"
)

type CheckPostDetail struct {
	Success bool
	Error   string
	Author  models.Author
	// Comments   []Comment // supposons que c'est le type retourné par GetComments
	NbrComment    int
	NbrLike       int
	LikeStatus    bool
	DisLikeStatus bool
}

// ______________________Handler
func PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := PostService.GetAllPosts()
		if err != nil {

		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			http.Error(w, "Failed to encode posts as JSON", http.StatusInternalServerError)
			return
		}
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	postValue := CheckPost(w, r)

	if !postValue.Success {
		http.Error(w, postValue.Error, http.StatusBadRequest)
		return
	}

	err := PostService.InsertPost(postValue)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// c'est une redirection temporaire
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}

func PostCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r)
	}
}

// ______________________fonction de traitement

func UploadImage(w http.ResponseWriter, r *http.Request) string {
	var photoURL string

	err := r.ParseMultipartForm(10 << 10)
	if err != nil {
		return "err400"
	}
	file, handler, err := r.FormFile("file")
	if file != nil {
		if err != nil {
			return "err400"
		}
		defer file.Close()
		if utils.IsValidImage(file, handler) {
			return "err400"
		}
		if handler.Size > 20<<20 {
			return "err408"
		}
		// path temporaire
		dirPath := "./web/static/upload"
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, 0755)
			if err != nil {
				return "err500"
			}
		}
		tempFile, err := os.CreateTemp(dirPath, "upload-*"+filepath.Ext(handler.Filename))
		if err != nil {
			return "err500"
		}
		defer tempFile.Close()
		_, err = io.Copy(tempFile, file)
		if err != nil {
			return "err500"
		}
		photoURL = tempFile.Name()
	}
	return photoURL
}

// ______________________fonction de verification
func CheckPost(w http.ResponseWriter, r *http.Request) models.CheckResult {
	content := strings.TrimSpace(r.FormValue("thread"))
	privacy := r.FormValue("privacy")
	photoURL := UploadImage(w, r)

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
