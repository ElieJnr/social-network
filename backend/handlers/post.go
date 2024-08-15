package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"strings"
	"time"
	"unicode/utf8"
)

type CheckResult struct {
	Success      bool
	Error        string
	PhotoURL     string
	Content      string
	Status       string
	AllowedUsers []string 
}

type CheckPostDetail struct {
    Success    bool
    Error      string
    Author     string
    // Comments   []Comment // supposons que c'est le type retourné par GetComments
    NbrComment int
    NbrLike    int
    LikeStatus   bool 
	DisLikeStatus bool

}
// __________________ Handler
func PostHandler(db *sqlite.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logique pour gérer les utilisateurs
		// users := []string{"user1", "user2", "user3"}
		// json.NewEncoder(w).Encode(users)
	}
}

func PostCreateHandler(db *sqlite.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreatePost(w, r, db)
	}
}

// ________________fonction de traitement
func CreatePost(w http.ResponseWriter, r *http.Request, db *sqlite.DB) {
	postValue := CheckPost(w, r)

	if !postValue.Success {
		http.Error(w, postValue.Error, http.StatusBadRequest)
		return
	}

	err := InsertPost(postValue, db)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// c'est une redirection temporaire
	http.Redirect(w, r, "/post", http.StatusSeeOther)
}

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
		if IsValidImage(file, handler) {
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

func GetAllPosts() ([]models.Posts, error) {
    query := `SELECT post_id, user_id, PhotoURL, content, status, creation_date FROM Posts ORDER BY creation_date DESC`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var allPosts []models.Posts
    for rows.Next() {
        var post models.Posts
        err := rows.Scan(&post.PostID, &post.UserID, &post.Image_url, &post.Content, &post.Post_status, &post.Creation_date)
        if err != nil {
            fmt.Println("err: rowsScan", err)
            continue
        }

        postDetail := CheckPostInformation(post.PostID, post.UserID)
        if !postDetail.Success {
            fmt.Println(postDetail.Error)
            continue
        }

        // Ajoutez les informations du postDetail à votre structure post

        post.Formated_date = FormatTimeAgo(post.Creation_date)
        allPosts = append(allPosts, post)
    }
    return allPosts, nil
}

func GetComments() {
}

// _________________fonction de verification
func CheckPost(w http.ResponseWriter, r *http.Request) CheckResult {
	content := strings.TrimSpace(r.FormValue("thread"))
	privacy := r.FormValue("privacy")
	photoURL := UploadImage(w, r)

	var allowedUsers []string
	if privacy == "almost_private" {
		allowedUsers = r.Form["allowedUsers"] // doit contenir les Userid des utilisateurs autorisés à voir le post
	}

	isValid, validationError := IsValidPost(content, privacy, photoURL, allowedUsers)
	if !isValid {
		return CheckResult{
			Success: false,
			Error:   validationError,
		}
	}

	return CheckResult{
		Success:  true,
		PhotoURL: photoURL,
		Content:  content,
		Status:   privacy,
		AllowedUsers: allowedUsers,
	}
}

func IsValidPost(content string, privacy string, photoURL string, allowedUsers []string) (bool, string) {
	if len(content) == 0 || utf8.RuneCountInString(content) > 500 {
		return false, "Bad Request: Invalid content length"
	}

	if privacy != "public" && privacy != "private" && privacy != "almost_private" {
		return false, "Bad Request: Invalid privacy setting"
	}

	if privacy == "almost_private" && len(allowedUsers) == 0 {
		return false, "Bad Request: No users specified for almost_private post"
	}

	if photoURL != "" {
		if photoURL == "err400" {
			return false, "Bad Request: Invalid image upload"
		} else if photoURL == "err500" {
			return false, "Internal Server Error: Image upload failed"
		}
	}

	return true, ""
}

func IsValidImage(file multipart.File, handler *multipart.FileHeader) bool {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return false
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buff)

	switch contentType {
	case "image/svg+xml", "image/jpeg", "image/gif", "image/png":
	default:
		fmt.Println(contentType)
		fmt.Println("type")
		return false
	}
	return true
}

func CheckPostInformation(postID, userID int) CheckPostDetail {
    author, err := GetPostAuthor(userID)
    if err != nil {
        return CheckPostDetail{
            Success: false,
            Error:   fmt.Sprintf("Err: GetPostAuthor: %v", err),
        }
    }

    comments, err := GetComments(postID)
    if err != nil {
        return CheckPostDetail{
            Success: false,
            Error:   fmt.Sprintf("Err: GetComments: %v", err),
        }
    }

    nbrComment, err := GetNbrComment(postID)
    if err != nil {
        return CheckPostDetail{
            Success: false,
            Error:   fmt.Sprintf("Err: GetNbrComment: %v", err),
        }
    }

    nbrLike, err := GetNbrLike(postID)
    if err != nil {
        return CheckPostDetail{
            Success: false,
            Error:   fmt.Sprintf("Err: GetNbrLike: %v", err),
        }
    }

    LikeStatus, err := GetLikeDislikeStatus(postID, userID)
    if err != nil {
        return CheckPostDetail{
            Success: false,
            Error:   fmt.Sprintf("Err: GetStatusLike: %v", err),
        }
    }

    return CheckPostDetail{
        Success:    true,
        Author:     author,
        // Comments:   comments,
        NbrComment: nbrComment,
        NbrLike:    nbrLike,
        LikeStatus:   LikeStatus,
    }
}



// _________________fonction DB
// recuperer l'id de l'utilisateur courant
func GetCurrentUser() (int, error) {
	return 1, nil
}

func GetNbrComment(postId int) {
}

func GetNbrLike(postId int) {
}

func GetNbrDislike(postId int) {
}

func GetLikeDislikeStatus(postId, userId int) {
}

func GetPostAuthor(userId int) {
}

// _________________fonction d'inserstion
func InsertPost(postValue CheckResult, db *sqlite.DB) error {
	userId, err := GetCurrentUser()
	if err != nil {
		return err
	}

	sqlDB := db.GetDB()
	tx, err := sqlDB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insérer le post principal
	result, err := tx.Exec("INSERT INTO Posts (user_id, content, image_url, statut) VALUES (?, ?, ?, ?)", 
		userId, postValue.Content, postValue.PhotoURL, postValue.Status)
	if err != nil {
		return err
	}

	lastInsertPostID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Préparer la requête d'insertion pour UserPosts
	stmt, err := tx.Prepare("INSERT INTO UserPosts (post_id, user_id, statut) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insérer dans UserPosts
	if postValue.Status == "public" || postValue.Status == "private" {
		_, err = stmt.Exec(nil, lastInsertPostID, postValue.Status)
		if err != nil {
			return err
		}
	} else if postValue.Status == "almost_private" {
		// Insérer pour le créateur du post
		_, err = stmt.Exec(lastInsertPostID, userId, postValue.Status)
		if err != nil {
			return err
		}

		// Insérer pour chaque utilisateur autorisé
		for _, allowedUserID := range postValue.AllowedUsers {
			_, err = stmt.Exec(lastInsertPostID, allowedUserID, postValue.Status)
			if err != nil {
				return err
			}
		}
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}

//__________________fonction utilitaire
func FormatTimeAgo(creationDate time.Time) string {
	now := time.Now()
	diff := now.Sub(creationDate)

	var res string

	switch {
	case diff.Hours() >= 24*365:
		years := int(diff.Hours() / (24 * 365))
		res = fmt.Sprintf("%d year", years)
		if years > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24*30:
		months := int(diff.Hours() / (24 * 30))
		res = fmt.Sprintf("%d month", months)
		if months > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24*7:
		weeks := int(diff.Hours() / (24 * 7))
		res = fmt.Sprintf("%d week", weeks)
		if weeks > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 24:
		days := int(diff.Hours() / 24)
		res = fmt.Sprintf("%d day", days)
		if days > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Hours() >= 1:
		hours := int(diff.Hours())
		res = fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			res += "s"
		}
		res += " ago"
	case diff.Minutes() >= 1:
		minutes := int(diff.Minutes())
		res = fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			res += "s"
		}
		res += " ago"
	default:
		seconds := int(diff.Seconds())
		res = fmt.Sprintf("%d second", seconds)
		if seconds != 1 {
			res += "s"
		}
		res += " ago"
	}

	return res
}
