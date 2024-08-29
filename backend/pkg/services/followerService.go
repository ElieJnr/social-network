package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"

	"github.com/gofrs/uuid/v5"
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

func (f *FollowerService) FollowUserOrUpdateStatus(userId uuid.UUID, followedUser uuid.UUID, statut bool) error {
	var existingId string
	queryCheck := `
		SELECT id 
		FROM Followers 
		WHERE userId = ? AND followedId = ?
	`
	err := f.GetDB().QueryRow(queryCheck, userId, followedUser).Scan(&existingId)

	if err != nil {
		if err == sql.ErrNoRows {
			id, err := utils.GenerateUuid()
			if err != nil {
				return fmt.Errorf("could not generate UUID: %w", err)
			}
			queryInsert := `
				INSERT INTO Followers (id, userId, followedId, statut) 
				VALUES (?, ?, ?, ?)
			`
			_, err = f.GetDB().Exec(queryInsert, id, userId, followedUser, statut)
			if err != nil {
				return fmt.Errorf("could not insert follower: %w", err)
			}
		} else {
			return fmt.Errorf("error checking existing follower: %w", err)
		}
	} else {
		queryUpdate := `
			UPDATE Followers 
			SET statut = ? 
			WHERE id = ?
		`
		_, err = f.GetDB().Exec(queryUpdate, statut, existingId)
		if err != nil {
			return fmt.Errorf("could not update follower status: %w", err)
		}
	}
	return nil
}

func (f *FollowerService) GetUserFollow(userId string, ok bool) ([]models.Follower, error) {
	var follows []models.Follower

	query := `
    SELECT f.id, f.userId, f.followedId, f.statut
    FROM Followers f
    INNER JOIN Users u ON f.userId = u.id
    WHERE f.followedId = ? AND statut = TRUE`

	if !ok {
		query = `
    SELECT f.id, f.userId, f.followedId, f.statut
    FROM Followers f
    INNER JOIN Users u ON f.userId = u.id
    WHERE f.userId = ? AND statut = TRUE`
	}

	rows, err := f.GetDB().Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var follow models.Follower
		if err := rows.Scan(&follow.Id, &follow.UserId, &follow.FollowedId, &follow.Statut); err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		follows = append(follows, follow)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return follows, nil
}
