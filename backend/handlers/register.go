package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"socialNetwork/pkg/models"
	"socialNetwork/utils"

	"golang.org/x/crypto/bcrypt"
)

// Response structure
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func RegistrationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse multipart form, with a maximum of 10MB for uploaded files
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			fmt.Println("Could not parse form:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusBadRequest,
				Message: "Could not parse form",
			})
			return
		}

		var newUser models.User
		newUser.Email = r.FormValue("email")
		newUser.Password = r.FormValue("password")
		newUser.Firstname = r.FormValue("firstname")
		newUser.Lastname = r.FormValue("lastname")
		newUser.Username = r.FormValue("username")
		newUser.Bio = r.FormValue("bio")
		newUser.DateOfBirth = r.FormValue("dateOfBirth")

		imgPath, err := utils.UploadImage(w, r, "register")
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		newUser.Avatar = imgPath

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Failed to hash password:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to hash password",
			})
			return
		}

		newUser.Password = string(hashedPassword)

		emailExists, err := UserService.EmailExists(newUser.Email)
		if err != nil {
			fmt.Println("Failed to check if email exists:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if emailExists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusMethodNotAllowed,
				Message: "Email already exists",
			})
			return
		}

		usernameExists, err := UserService.UsernameExists(newUser.Username)
		if err != nil {
			fmt.Println("Failed to check if username exists:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		if usernameExists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusMethodNotAllowed,
				Message: "Username already exists",
			})
			return
		}

		err = UserService.CreateUser(newUser)
		if err != nil {
			fmt.Println("Failed to create user:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusCreated,
			Message: "User created successfully",
			Data:    newUser,
		})
	}
}
