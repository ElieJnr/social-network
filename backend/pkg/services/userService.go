package services

import (
	"database/sql"
	"fmt"
	"net/http"
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
	id, err := uuid.NewV4()
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

func (u *UserService) UpdateUser(private bool, userId uuid.UUID) error {
	queryUpdate := `
 			UPDATE Users 
 			SET isPrivate = ? 
 			WHERE id = ?
 		`
	_, err := u.GetDB().Exec(queryUpdate, private, userId)
	if err != nil {
		fmt.Println("Error: cannot update user", err)
		return fmt.Errorf("could not update follower status: %w", err)
	}
	return nil
}

func (u *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) ([]models.User, error) {
	query := `SELECT id, email, firstname, lastname, dateOfBirth, avatar, username, bio, isPrivate FROM Users`
	rows, err := u.GetDB().Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve users: %w", err)
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Username, &user.Bio, &user.IsPrivate)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %w", err)
		}
		followService := NewFollowerService()
		postService := NewPostService()
		followers, err := followService.GetUserFollow(user.Id, true)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %w", err)
		}

		follows, err := followService.GetUserFollow(user.Id, false)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %w", err)
		}
		posts, err := postService.GetPostsById(w,r,string(user.Id))
		if err != nil {
			return nil, fmt.Errorf("could not scan posts: %w", err)
		}
		user.Followers = followers
		user.Follows = follows
		user.Posts = posts
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return users, nil
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

	if username == "" {
		return false, nil
	}

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

func (u *UserService) GetUserById(w http.ResponseWriter, r *http.Request, userId uuid.UUID) (models.User, error) {
	var user models.User
	err := u.GetDB().QueryRow("SELECT id, email, password, firstname, lastname, username, dateOfBirth, bio, avatar, isPrivate FROM Users WHERE id = ?", userId).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Username, &user.DateOfBirth, &user.Bio, &user.Avatar, &user.IsPrivate)
	if err != nil {
		return models.User{}, err
	}
	followService := NewFollowerService()
	postService := NewPostService()
	followers, err := followService.GetUserFollow(user.Id, true)
	if err != nil {
		return models.User{}, fmt.Errorf("could not scan user: %w", err)
	}

	follows, err := followService.GetUserFollow(user.Id, false)
	if err != nil {
		return models.User{}, fmt.Errorf("could not scan user: %w", err)
	}
	requestFollow, err := followService.GetUserRequest(user.Id)
	if err != nil {
		return models.User{}, fmt.Errorf("could not scan requestFollow: %w", err)
	}
	posts, err := postService.GetPostsById(w,r,string(user.Id))
	if err != nil {
		return models.User{}, fmt.Errorf("could not scan posts: %w", err)
	}

	user.RequestF = requestFollow
	user.Followers = followers
	user.Follows = follows
	user.Posts = posts
	return user, nil
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

func (u *UserService) DeleteColumnByID(tableName string, userID string) error {
	// Préparer la requête SQL pour supprimer une ligne
	query := fmt.Sprintf("DELETE FROM %s WHERE userId = ?", tableName)

	// Créer la requête préparée
	stmt, err := u.GetDB().Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Exécuter la requête avec l'ID de l'utilisateur
	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	fmt.Printf("User with ID %s has been deleted from table %s\n", userID, tableName)
	return nil
}
