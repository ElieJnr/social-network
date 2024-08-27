package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"github.com/gofrs/uuid/v5"
)

type UserService struct {
	db *sql.DB
}

func NewUserService() *UserService {

	dbs := sqlite.GlobalDB

	return &UserService{
		db: dbs.GetDB(),
	}
}

func (u *UserService) GetDB() *sql.DB {
	return u.db
}

func (u *UserService) SetDB(db *sql.DB) {
	u.db = db
}

func (u *UserService) CreateUser(user models.User) error {
	// Préparer la requête d'insertion
	id,  err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}
	query := `
		INSERT INTO Users (id,email, password, firstname, lastname, dateOfBirth, avatar, username, bio, isPrivate) 
		VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = u.GetDB().Exec(query, id, user.Email, user.Password, user.Firstname, user.Lastname, user.DateOfBirth, user.Avatar, user.Username, user.Bio, false)
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}

	return nil
}


func (u *UserService) UserExists(EmailOrName string) (*models.User, error) {
	var user models.User

	err := u.GetDB().QueryRow("SELECT id,username,password,email FROM Users WHERE username = ? OR email = ?", EmailOrName, EmailOrName).Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserService) UsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ? LIMIT 1)"
	err := u.GetDB().QueryRow(query, username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

func (u *UserService) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ? LIMIT 1)"
	err := u.GetDB().QueryRow(query, email).Scan(&exists)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
