package handlers

import (
	"encoding/json"
	"net/http"

	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var newUser models.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		session := uuid.New()
		newUser.Session = session.String()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		newUser.Password = string(hashedPassword)
		userService := services.NewUserService()
		err = userService.CreateUser(newUser.Email, newUser.Password, newUser.Firstname, newUser.Lastname, newUser.DateOfBirth, newUser.Avatar, newUser.Username, newUser.Bio, newUser.Session)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("User created successfully")
	}
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var credentials struct {
			EmailOrUsername string `json:"emailOrUsername"`
			Password        string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService := services.NewUserService()

		pass, id, err := userService.UserExists(credentials.EmailOrUsername)
		if err != nil {
			http.Error(w, "Invalid credentiale", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(credentials.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Générez un token JWT ou une autre méthode pour maintenir la session
		session := uuid.New()

		userService.UpdateSessionByID(id, session.String())

		json.NewEncoder(w).Encode(map[string]string{"token": session.String()})

	}
}
