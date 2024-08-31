package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/utils"
)

type LikeService struct {
	db *sql.DB
}

func NewLikeService() *LikeService {
	dbs := sqlite.GlobalDB
	return &LikeService{
		db: dbs.GetDB(),
	}
}

func (l *LikeService) LikeTreatment(postID string, w http.ResponseWriter, r *http.Request) error {

	currentUser, err := utils.CurrentUser(w, r)
	if err != nil {
		fmt.Println("failed to get current user")
		return fmt.Errorf("failed to get current user: %w", err)
	}

	// Vérification du statut actuel du like
	liked, err := l.GetExistingLike(postID, currentUser.UserId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error checking like status")
		return fmt.Errorf("error checking like status: %w", err)
	}

	if err == sql.ErrNoRows {
		// Pas de like existant, insérer un nouveau like
		err = l.InsertLike(postID, currentUser.UserId, true)
		if err != nil {
			fmt.Println("error inserting like")
			return fmt.Errorf("error inserting like: %w", err)
		}
	} else {
		// Basculer le statut du like
		// newLikeStatus := !liked
		err = l.UpdateLike(postID, currentUser.UserId, !liked)
		if err != nil {
			fmt.Println("error updating like")
			return fmt.Errorf("error updating like: %w", err)
		}
	}

	// fmt.Println("Like inserted/updated successfully")
	return nil
}

func (l *LikeService) GetExistingLike(postID, userID string) (bool, error) {
	query := `SELECT liked FROM LikesDislikes WHERE postId = ? AND userId = ?`
	var liked bool
	row := l.db.QueryRow(query, postID, userID)
	err := row.Scan(&liked)
	if err != nil {
		return false, err
	}
	return liked, nil
}

func (l *LikeService) InsertLike(postID, userID string, liked bool) error {
	// fmt.Println("InsertLike here we GOOOOOOOOOOO")
	query := `INSERT INTO LikesDislikes (postId, userId, liked) VALUES (?, ?, ?)`
	_, err := l.db.Exec(query, postID, userID, liked)
	return err
}

func (l *LikeService) UpdateLike(postID, userID string, liked bool) error {
	query := `UPDATE LikesDislikes SET liked = ? WHERE postId = ? AND userId = ?`
	_, err := l.db.Exec(query, liked, postID, userID)
	return err
}
