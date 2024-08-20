package services

import (
	"database/sql"
	"errors"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
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

func (u *UserService) CreateUser(email, password, firstname, lastname, dateOfBirth, avatar, username, bio, session string) error {
	// Vérifier si l'email existe déjà
	emailExists, err := u.EmailExists(email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}

	// Vérifier si le nom d'utilisateur existe déjà (s'il est fourni)
	if username != "" {
		usernameExists, err := u.UsernameExists(username)
		if err != nil {
			return err
		}
		if usernameExists {
			return errors.New("username already exists")
		}
	}

	// Préparer la requête d'insertion
	query := `
		INSERT INTO Users (id,email, passwords, firstname, lastname, dateOfBirth, avatar, username, bio, isPrivate) 
		VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = u.GetDB().Exec(query, session, email, password, firstname, lastname, dateOfBirth, avatar, username, bio, false)
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}

	return nil
}

func (u *UserService) UserExists(identifier string) (string, string, error) {
	var password string
	var id string
	err := u.GetDB().QueryRow("SELECT passwords, id FROM Users WHERE username = ? OR email = ?", identifier, identifier).Scan(&password, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", errors.New("user not found")
		}
		return "", "", err
	}
	return password, id, nil
}

func (u *UserService) UsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ? LIMIT 1)"
	err := u.GetDB().QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserService) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = ? LIMIT 1)"
	err := u.GetDB().QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserService) UpdateSessionByID(session, id string) error {
	// Préparation de la requête SQL pour mettre à jour la session
	query := "UPDATE Users SET session = ? WHERE id = ?"

	// Exécution de la requête préparée
	stmt, err := u.GetDB().Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Lier les valeurs à la requête et l'exécuter
	_, err = stmt.Exec(session, id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
