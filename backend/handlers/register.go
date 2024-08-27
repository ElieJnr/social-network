package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"socialNetwork/pkg/models"
	"socialNetwork/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegistrationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")

		if r.Method != http.MethodPost {
			fmt.Println("1")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse multipart form, with a maximum of 10MB for uploaded files
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			fmt.Println("2")
			http.Error(w, "Could not parse form", http.StatusBadRequest)
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

		imgPath := utils.UploadImage(w, r, "register")

		fmt.Println("imgpath",imgPath)
		newUser.Avatar = imgPath

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("3")
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		newUser.Password = string(hashedPassword)

		fmt.Println("4")

		emailExists, err := UserService.EmailExists(newUser.Email)
		if err != nil {
			fmt.Println("5")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if emailExists {
			fmt.Println("6")
			http.Error(w, "Email already exists", http.StatusMethodNotAllowed)
			return
		}

		usernameExists, err := UserService.UsernameExists(newUser.Username)
		if err != nil {
			fmt.Println("8")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if usernameExists {
			fmt.Println("9")
			http.Error(w, "Username already exists", http.StatusMethodNotAllowed)
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
