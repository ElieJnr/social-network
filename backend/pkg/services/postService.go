package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"

	"github.com/google/uuid"
)

type PostService struct {
	db *sql.DB
}

func NewPostService() *PostService {
	dbs := sqlite.GlobalDB
	return &PostService{
		db: dbs.GetDB(),
	}
}

func (p *PostService) GetDB() *sql.DB {
	return p.db
}

func (p *PostService) SetDB(db *sql.DB) {
	p.db = db
}

// _______________________foonction d'insertion
func (p *PostService) InsertPost(postValue models.CheckResult) error {
	// récuperation fictif en attendant qu'on ai une veritable fonction pour connitre qui se connecte en ce moment
	userId, err := GetCurrentUser()
	if err != nil {
		return err
	}

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO Posts (userId, content, imageUrl, statut) VALUES (?, ?, ?, ?)",
		userId, postValue.Content, postValue.PhotoURL, postValue.Status)
	if err != nil {
		return err
	}

	lastInsertPostID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO UserPosts (postId, userId, statut) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if postValue.Status == "public" || postValue.Status == "private" {
		_, err = stmt.Exec(lastInsertPostID, userId, postValue.Status)
		if err != nil {
			return err
		}
	} else if postValue.Status == "almost_private" {
		_, err = stmt.Exec(lastInsertPostID, userId, postValue.Status)
		if err != nil {
			return err
		}

		for _, allowedUserID := range postValue.AllowedUsers {
			_, err = stmt.Exec(lastInsertPostID, allowedUserID, postValue.Status)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

// ______________________fonction de recuperation
func (p *PostService) GetAllPosts() ([]models.Posts, error) {
	query := `SELECT id, userId, content, imageUrl, statut, createDate FROM Posts ORDER BY createDate DESC`
	rows, err := p.db.Query(query) // Utilisation de p.db au lieu de db
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

		if err := postDetails(p.db, &post); err != nil {
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
	// TODO: Implement this function
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

	isVisible, err := CheckVisibility(db, post.UserID, post.PostID, currentUserId, post.Post_status)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}

	post.Author = author
	post.Formated_date = utils.FormatTimeAgo(post.Creation_date)
	post.Can_see = isVisible
	post.Like_nbr = nbrLike
	post.Comments_nbr = nbrComment
	post.Like_status = likeStatus
	post.Dislike_nbr = nbrDislike
	post.Dislike_status = dislikeStatus

	return nil
}

func CheckVisibility(db *sql.DB, userId uuid.UUID, postId uuid.UUID, currentUserId uuid.UUID, postPrivacy string) (bool, error) {
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
		query := `SELECT COUNT(*) FROM Followers WHERE userId = ? AND followerId = ?`
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
		query := `SELECT COUNT(*) FROM UserPost WHERE postId = ? AND userId = ?`
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

// fictif
func GetCurrentUser() (uuid.UUID, error) {
	return uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), nil
}

func GetNbrComment(db *sql.DB, postId uuid.UUID) (int, error) {
	return getCountForPost(db, "Comments", "AND postid = ?", postId)
}

func GetNbrLike(db *sql.DB, postId uuid.UUID) (int, error) {
	return getCountForPost(db, "LikesDislikes", "AND liked = TRUE", postId)
}

func GetNbrDislike(db *sql.DB, postId uuid.UUID) (int, error) {
	return getCountForPost(db, "LikesDislikes", "AND disliked = TRUE", postId)
}

func GetLikeDislikeStatus(db *sql.DB, postId uuid.UUID, userId uuid.UUID) (bool, bool, error) {
	query := `
        SELECT liked, disliked 
        FROM LikesDislikes 
        WHERE postId = ? AND userId = ?
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

func GetPostAuthor(db *sql.DB, userId uuid.UUID) (models.Author, error) {

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

func getCountForPost(db *sql.DB, tableName string, condition string, postId uuid.UUID) (int, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE postId = ? %s`, tableName, condition)

	var count int
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error querying database: %v", err)
	}
	return count, nil
}
