package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type GroupeService struct {
	db *sql.DB
}

func NewGroupeService() *GroupeService {
	dbs := sqlite.GlobalDB

	return &GroupeService{
		db: dbs.GetDB(),
	}
}

func (g *GroupeService) GetDB() *sql.DB {
	return g.db
}

func (g *GroupeService) CreateGroup(group models.Group) error {
	id, er := utils.GenerateUuid()

	if er != nil {
		return er
	}

	query := `INSERT INTO Groups (id,title,description, userId)
		VALUES (?,?,?,?)`

	_, err := g.GetDB().Exec(query, id, group.Title, group.Description, group.UserId)

	if err != nil {
		return fmt.Errorf("failed to create group: %w", err)
	}

	return nil
}

func (g *GroupeService) GetGroups() ([]models.Group, error) {
	query := `SELECT * FROM Groups ORDER BY created_at DESC`
	fmt.Println("ici")
	rows, err := g.GetDB().Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		if err = rows.Scan(&group.Id, &group.Title, &group.Description, &group.UserId, &group.CreateAt); err != nil {
			return nil, fmt.Errorf("failed to scan group: %w", err)
		}
		fmt.Println(group)
		groups = append(groups, group)
	}
	return groups, nil
}
