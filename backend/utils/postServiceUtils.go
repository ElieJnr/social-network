package utils

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/models"
)

// se charge de vérifier si l'utiiateur peut voir le post
func CheckVisibility(db *sql.DB, userId string, postId string, currentUserId string, postPrivacy string, groupId string) (bool, error) {
	// Si le post est public, il est visible pour tout le monde
	if postPrivacy == "public" {
		return true, nil
	}

	// Si le post est privé, vérifier si l'utilisateur est l'auteur du post ou un follower
	if postPrivacy == "private" {
		return IsFollowing(db, userId, currentUserId)
	}

	// Si le post est "almost_private", vérifier si l'utilisateur est autorisé à voir le post
	if postPrivacy == "almost-private" {
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

	// Si le post est "group", vérifier si l'utilisateur est membre du groupe
	if postPrivacy == "group" {
		query := `SELECT COUNT(*) FROM Membership WHERE userId = ? AND groupId = ?`
		var count int
		err := db.QueryRow(query, currentUserId, groupId).Scan(&count)
		if err != nil {
			return false, fmt.Errorf("error checking group membership: %v", err)
		}

		// Si l'utilisateur est membre du groupe, retourner true
		if count > 0 {
			return true, nil
		}

		// Sinon, retourner false car l'utilisateur n'est pas membre du groupe
		return false, nil
	}

	// Par défaut, retourner false (post non visible)
	return false, nil
}

// Récupère le nombre de commentaires d'un post
func GetNbrComment(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "Comments", "postid = ?", postId)
}

// Récupère le nombre de likes d'un post
func GetNbrLike(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "LikesDislikes", "postid = ? AND liked = TRUE", postId)
}

// Récupère le nombre de dislikes d'un post
func GetNbrDislike(db *sql.DB, postId string) (int, error) {
	return getCountForPost(db, "LikesDislikes", "postid = ? AND disliked = TRUE", postId)
}

// Fonction utilitaire pour compter les enregistrements dans une table donnée avec une condition spécifique
func getCountForPost(db *sql.DB, tableName string, condition string, postId string) (int, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE %s`, tableName, condition)

	var count int
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("erreur lors de la requête SQL : %v", err)
	}
	return count, nil
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

// retourne tous les posts d'un user
func GetPostsById(db *sql.DB, userId string) ([]models.Posts, error) {
	query := `SELECT id, userId, content, imageUrl, statut, createDate FROM Posts WHERE userId = ? ORDER BY createDate DESC`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var ownPosts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Content, &post.Image_url, &post.Post_status, &post.Creation_date); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		ownPosts = append(ownPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return ownPosts, nil
}

// verifie si l'utilisateur est un follower pour les posts

func IsFollowing(db *sql.DB, userId string, currentUserId string) (bool, error) {
	if currentUserId == userId {
		return true, nil
	}

	// Vérifier si currentUserId est un ami de l'utilisateur
	query := `SELECT COUNT(*) FROM Followers WHERE userId = ? AND followedId = ? AND statut = true`
	var count int
	err := db.QueryRow(query, currentUserId, userId).Scan(&count)
	if err != nil {
		fmt.Println("error checking friend status")
		return false, err
	}

	// Si currentUserId est un ami, retourner true
	if count > 0 {
		return true, nil
	}

	// Sinon, retourner false car l'utilisateur n'est ni l'auteur ni un ami
	return false, nil

}
