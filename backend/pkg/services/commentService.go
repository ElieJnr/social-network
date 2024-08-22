package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type CommentService struct {
	db *sql.DB
}

func NewCommentService() *CommentService {
	dbs := sqlite.GlobalDB
	return &CommentService{
		db: dbs.GetDB(),
	}
}

func (c *CommentService) GetDB() *sql.DB {
	return c.db
}

func (c *CommentService) SetDB(db *sql.DB) {
	c.db = db
}

func (c *CommentService) InsertComment(commentValue models.CheckResult, w http.ResponseWriter, r *http.Request, postId string) error {
	user, err := utils.CurrentUser(w, r)
	if err != nil {
		return err
	}

	commentID, err := utils.GenerateUuid()
	if err != nil {
		return err
	}

	_, err = c.db.Exec("INSERT INTO Comments (id, postId, userId, content, imageUrl) VALUES (?, ?, ?, ?, ?)",
		commentID, postId, user.UserId, commentValue.Content, commentValue.PhotoURL)
	if err != nil {
		return err
	}

	return nil
}

func GetComments(db *sql.DB, postId string) ([]models.Comment, error) {
	query := `SELECT id, userId, content, imageUrl, createDate FROM Comments WHERE postId = ? ORDER BY createDate DESC`
	rows, err := db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.Content, &comment.Image_url, &comment.Creation_date); err != nil {
			return nil, fmt.Errorf("failed to scan comment row: %w", err)
		}

		author, err := utils.GetAuthor(db, comment.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to get post author: %w", err)
		}

		comment.Author = author
		comment.Formated_date = utils.FormatTimeAgo(comment.Creation_date)
		comments = append(comments, comment)
	}

	return comments, nil
}
