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
	dbs, err := sqlite.NewDatabase()
	if err != nil {
		fmt.Println(err)
	}
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
		INSERT INTO Users (email, password, firstname, lastname, date_of_birth, avatar, username, bio, isPrivate, session) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = u.GetDB().Exec(query, email, password, firstname, lastname, dateOfBirth, avatar, username, bio, false, session)
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}

	return nil
}

// func (u *UserService) CreateUser(username string, age, genre, firstname, lastname, email, password string) error {
// }

// func (u *UserService) GetAllUsers() ([]models.User, error) {
// 	rows, err := u.GetDB().Query(`SELECT id, username, age, genre, firstname, lastname, email, password FROM users`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// }

// func (u *UserService) GetUserById(id uuid.UUID) (models.User, error) {

// 	err := u.db.QueryRow(`SELECT id, username, age, genre, firstname, lastname, email, password FROM users WHERE id = ?`, id.String()).Scan(&idStr, &username, &age, &genre, &firstname, &lastname, &email, &password)
// 	if err != nil {
// 		//fmt.Println(id)
// 		fmt.Println(err)
// 		return models.User{}, err

// }
// }

// func (u *UserService) GetUserByUsernameOREmail(username string) (models.User, error) {

// 	// err := u.db.QueryRow(`SELECT id, username, age, genre, firstname, lastname, email, password FROM users WHERE username = ? OR email = ?`, username, username).Scan(&id, &userName, &age, &genre, &firstname, &lastname, &email, &password)
// 	// if err != nil {

// 	// 	return models.User{}, err
// 	// }

// }

func (u *UserService) UserExists(identifier string) (string, int, error) {
	var password string
	var id int
	err := u.GetDB().QueryRow("SELECT password, id FROM users WHERE username = ? OR email = ?", identifier, identifier).Scan(&password, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", -1, errors.New("user not found")
		}
		return "", -1, err
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

func (u *UserService) UpdateSessionByID(id int, session string) error {
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