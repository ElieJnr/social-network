package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type FollowerService struct {
	db *sql.DB
}

func NewFollowerService() *FollowerService {
	dbs := sqlite.GlobalDB
	return &FollowerService{
		db: dbs.GetDB(),
	}
}

func (f *FollowerService) GetDB() *sql.DB {
	return f.db
}

func (f *FollowerService) SetDB(db *sql.DB) {
	f.db = db
}

func (f *FollowerService) FollowUser(userId string, followedUser []models.User) error {
	for _, user := range followedUser {
		id, e := utils.GenerateUuid()
		if e != nil {
			return e
		}
		query := `
			INSERT INTO Followers (id, userId, followedId) 
			VALUES (?, ?, ?)

		`
		_, err := f.GetDB().Exec(query, id, userId, user.Id)
		if err != nil {
			return fmt.Errorf("could not insert follower: %w", err)
		}
	}
	return nil
}

func (f *FollowerService) GetUserFollow(userId string, ok bool) ([]models.User, error) {
	var users []models.User

	query := `
        SELECT u.id, u.email, u.passwords, u.firstname, u.lastname, u.username, u.dateOfBirth, u.bio, u.avatar, u.isPrivate 
        FROM Users u
        INNER JOIN Followers f ON u.id = f.followedId
        WHERE f.userId = ?`
	if !ok {
		query = `
        SELECT u.id, u.email, u.passwords, u.firstname, u.lastname, u.username, u.dateOfBirth, u.bio, u.avatar, u.isPrivate 
        FROM Users u
        INNER JOIN Followers f ON u.id = f.userId
        WHERE f.followedId = ?`
	}

	rows, err := f.GetDB().Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Username, &user.DateOfBirth, &user.Bio, &user.Avatar, &user.IsPrivate); err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return users, nil
}
