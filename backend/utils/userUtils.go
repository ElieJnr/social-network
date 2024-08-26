package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
)

type contextKey string

const UserContextKey contextKey = "user"

func CurrentUser(w http.ResponseWriter, r *http.Request) (*models.Session, error) {
	user, ok := r.Context().Value(UserContextKey).(*models.Session)
	fmt.Println("user", user)
	if !ok || user == nil {
		// http.Error(w, "User not found in context", http.StatusUnauthorized)
		return nil, errors.New("user not found in context")
	}
	return user, nil
}

// se charge de récuperer l'auteur d'un post
func GetAuthor(db *sql.DB, userId string) (models.Author, error) {

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
			return author, fmt.Errorf("no user found with ID %s", userId)
		}
		return author, fmt.Errorf("error querying database: %v", err)
	}

	return author, nil
}
