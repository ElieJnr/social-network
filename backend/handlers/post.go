package handlers

import (
	"database/sql"
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
	Success bool
	Error   string
	Author  models.Author
	// Comments   []Comment // supposons que c'est le type retourné par GetComments
	NbrComment    int
	NbrLike       int
	LikeStatus    bool
	DisLikeStatus bool
}

//______________________Handler
func PostHandler(db *sqlite.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetAllPosts(w, r, db)
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

//______________________fonction de traitement
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

func postDetails(db *sql.DB, post *models.Posts) error {
	currentUserId, err := GetCurrentUser()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	author, err := GetPostAuthor(db, post.UserID)
	if err != nil {
		return fmt.Errorf("failed to get post author: %w", err)
	}
	
	nbrComment, err := GetNbrComment(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get number of comments: %w", err)
	}
	
	nbrLike, err := GetNbrLike(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get number of likes: %w", err)
	}

	nbrDislike, err := GetNbrDislike(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get number of likes: %w", err)
	}
	
	likeStatus, dislikeStatus, err := GetLikeDislikeStatus(db, post.PostID, post.UserID)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}

	isVisible, err := CheckVisibility(db, post.UserID, post.PostID,currentUserId, post.Post_status)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}
	
	post.Author = author
	post.Formated_date = FormatTimeAgo(post.Creation_date)
	post.Can_see = isVisible
	post.Like_nbr = nbrLike
	post.Comments_nbr = nbrComment
	post.Like_status = likeStatus
	post.Dislike_nbr = nbrDislike
	post.Dislike_status = dislikeStatus

	return nil
}

//______________________fonction de verification
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
		Success:      true,
		PhotoURL:     photoURL,
		Content:      content,
		Status:       privacy,
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

func CheckVisibility(db *sql.DB, userId int, postId int, currentUserId int, postPrivacy string) (bool, error) {
	// Si le post est public, il est visible pour tout le monde
	if postPrivacy == "public" {
		return true, nil
	}

	// Si le post est privé, vérifier si l'utilisateur est l'auteur du post ou un follower
	if postPrivacy == "private" {
		if currentUserId == userId {
			return true, nil
		}

		// Vérifier si le currentUserId est un follower de l'auteur du post
		query := `SELECT COUNT(*) FROM Followers WHERE user_id = ? AND followerId = ?`
		var count int
		err := db.QueryRow(query, userId, currentUserId).Scan(&count)
		if err != nil {
			return false, fmt.Errorf("error checking follower status: %v", err)
		}

		// Si currentUserId est un follower, retourner true
		if count > 0 {
			return true, nil
		}

		// Sinon, retourner false car l'utilisateur n'est ni l'auteur ni un follower
		return false, nil
	}

	// Si le post est "almost_private", vérifier si l'utilisateur est autorisé à voir le post
	if postPrivacy == "almost_private" {
		query := `SELECT COUNT(*) FROM UserPost WHERE post_id = ? AND user_id = ?`
		var count int
		err := db.QueryRow(query, postId, currentUserId).Scan(&count)
		if err != nil {
			return false, fmt.Errorf("error checking almost_private access: %v", err)
		}

		// Si l'utilisateur est autorisé, retourner true
		if count > 0 {
			return true, nil
		}

		// Sinon, retourner false car l'utilisateur n'est pas autorisé
		return false, nil
	}

	// Par défaut, retourner false (post non visible)
	return false, nil
}


// ______________________fonction de recuperation
func GetAllPosts(w http.ResponseWriter, r *http.Request, sqlDB *sqlite.DB) ([]models.Posts, error) {
	db := sqlDB.GetDB()

	query := `SELECT post_id, user_id, content, image_url, statut, creatdate FROM Posts ORDER BY creatdate DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}
	defer rows.Close()

	var allPosts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Content, &post.Image_url, &post.Post_status, &post.Creation_date); err != nil {
			fmt.Println("err: rowsScan", err)
			continue
		}

		if err := postDetails(db, &post); err != nil {
			fmt.Println("err: postDetails", err)
			continue
		}

		allPosts = append(allPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over posts: %w", err)
	}

	return allPosts, nil
}

func GetComments() {
}

func GetCurrentUser() (int, error) {
	return 1, nil
}

func GetNbrComment(db *sql.DB, postId int) (int, error) {
    return getCountForPost(db, "Comments", "AND postid = ?", postId)
}

func GetNbrLike(db *sql.DB, postId int) (int, error) {
    return getCountForPost(db, "LikesDislikes", "AND liked = TRUE", postId)
}

func GetNbrDislike(db *sql.DB, postId int) (int, error) {
    return getCountForPost(db, "LikesDislikes", "AND disliked = TRUE", postId)
}

func GetLikeDislikeStatus(db *sql.DB, postId int, userId int) (bool, bool, error) {
    query := `
        SELECT liked, disliked 
        FROM LikesDislikes 
        WHERE post_id = ? AND user_id = ?
    `

    var liked, disliked bool
    err := db.QueryRow(query, postId, userId).Scan(&liked, &disliked)
    if err != nil {
        if err == sql.ErrNoRows {
            // If no row is found, the user hasn't liked or disliked the post
            return false, false, nil
        }
        return false, false, fmt.Errorf("error querying database: %v", err)
    }

    return liked, disliked, nil
}

func GetPostAuthor(db *sql.DB, userId int) (models.Author, error) {

	var author models.Author
	query := `
        SELECT firstname, lastname, username, avatar
        FROM Users
        WHERE id = ?
    `

	err := db.QueryRow(query, userId).Scan(
		&author.Firstname,
		&author.Lastname,
		&author.Username,
		&author.Avatar,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return author, fmt.Errorf("no user found with ID %d", userId)
		}
		return author, fmt.Errorf("error querying database: %v", err)
	}

	return author, nil
}

//______________________fonction d'inserstion
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

//______________________fonction utilitaire
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

func getCountForPost(db *sql.DB, tableName string, condition string, postId int) (int, error) {
    query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE post_id = ? %s`, tableName, condition)

    var count int
    err := db.QueryRow(query, postId).Scan(&count)
    if err != nil {
        return 0, fmt.Errorf("error querying database: %v", err)
    }
    return count, nil
}

