package handlers

import (
	"encoding/json"
	"net/http"

	"socialNetwork/pkg/models"

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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		newUser.Password = string(hashedPassword)

		// if Empty(newUser) {
		// 	http.Error(w, "Bad Request", http.StatusMethodNotAllowed)
		// 	return
		// }

		// Vérifier si l'email existe déjà
		emailExists, err := UserService.EmailExists(newUser.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if emailExists {
			http.Error(w, "Email already Exist", http.StatusMethodNotAllowed)
			return
		}

		// Vérifier si le nom d'utilisateur existe déjà (s'il est fourni)

		usernameExists, err := UserService.UsernameExists(newUser.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if usernameExists {
			http.Error(w, "Username already Exist", http.StatusMethodNotAllowed)
			return
		}

		err = UserService.CreateUser(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("User created successfully")
	}
}
