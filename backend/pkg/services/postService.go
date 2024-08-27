package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
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
func (p *PostService) InsertPost(postValue models.CheckResult, w http.ResponseWriter, r *http.Request) error {
	user, err := utils.CurrentUser(w, r)
	if err != nil {
		fmt.Println("error getting current user")
		return err
	}
	postID, err := utils.GenerateUuid()
	if err != nil {
		fmt.Println("error generating post ID")
		return err
	}

	tx, err := p.db.Begin()
	if err != nil {
		fmt.Println("error in transaction")
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO Posts (id, userId, content, imageUrl, statut) VALUES (?, ?, ?, ?, ?)",
		postID, user.UserId, postValue.Content, postValue.PhotoURL, postValue.Status)
	if err != nil {
		fmt.Println("error inserting post")
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO UserPost (postId, userId, statut) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("error preparing statement")
		return err
	}

	defer stmt.Close()

	if postValue.Status == "public" || postValue.Status == "private" {
		_, err = stmt.Exec(postID, user.UserId, postValue.Status)
		if err != nil {
			fmt.Println("error executing statement")
			return err
		}
	} else if postValue.Status == "almost_private" {
		_, err = stmt.Exec(postID, user.UserId, postValue.Status)
		if err != nil {
			fmt.Println("error executing statement")
			return err
		}

		for _, allowedUserID := range postValue.AllowedUsers {
			_, err = stmt.Exec(postID, allowedUserID, postValue.Status)
			if err != nil {
				fmt.Println("error executing statement")
				return err
			}
		}
	}

	fmt.Println("post inserted successfully")
	return tx.Commit()
}

// se charge de récuperer les posts et de le partager au service de post
func (p *PostService) GetAllPosts(w http.ResponseWriter, r *http.Request) ([]models.Posts, error) {
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

		if err := postDetails(p.db, &post, w, r); err != nil {
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

// se charge de récuperer les informations d'un post ie author, nbr de like, qui peut voir le post, nbr de commentaires...
func postDetails(db *sql.DB, post *models.Posts, w http.ResponseWriter, r *http.Request) error {
	currentUser, err := utils.CurrentUser(w, r)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	author, err := utils.GetAuthor(db, post.UserID)
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

	isVisible, err := CheckVisibility(db, post.UserID, post.PostID, currentUser.UserId, post.Post_status)
	if err != nil {
		return fmt.Errorf("failed to get like/dislike status: %w", err)
	}

	allComments, err := GetComments(db, post.PostID)
	if err != nil {
		return fmt.Errorf("failed to get comments: %w", err)
	}

	post.Author = author
	post.Formated_date = utils.FormatTimeAgo(post.Creation_date)
	post.Can_see = isVisible
	post.Comments = allComments
	post.Like_nbr = nbrLike
	post.Comments_nbr = nbrComment
	post.Like_status = likeStatus
	post.Dislike_nbr = nbrDislike
	post.Dislike_status = dislikeStatus

	return nil
}

// se charge de vérifier si l'utiiateur peut voir le post
func CheckVisibility(db *sql.DB, userId string, postId string, currentUserId string, postPrivacy string) (bool, error) {
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

// se charge de récuperer le nombre de commentaires d'un post
func GetNbrComment(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "Comments", "AND postid = ?", postId)
}

// se charge de récuperer le nombre de like d'un post
func GetNbrLike(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "LikesDislikes", "AND liked = TRUE", postId)
}

// se charge de récuperer le nombre de dislike d'un post
func GetNbrDislike(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "LikesDislikes", "AND disliked = TRUE", postId)
}

// se charge de récuperer le status de like et dislike d'un post
func GetLikeDislikeStatus(db *sql.DB, postId string, userId string) (bool, bool, error) {
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

// fonction creer pour eviter la redondance de code pour le nombre de like et dislike et peut etre utilisé pour tous ce qui nécessite un count
func getCountForPost(db *sql.DB, tableName string, condition string, postId string) (int, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE postId = ? %s`, tableName, condition)

	var count int
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error querying database: %v", err)
	}
	return count, nil
}
