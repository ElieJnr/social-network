package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var credentials struct {
			EmailOrName string
			Password    string
		}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			fmt.Println("bad")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(credentials.EmailOrName)
		user, e := UserService.UserExists(credentials.EmailOrName)
		// id := user.Id
		// username := user.Username
		//  := user.Password
		// email := user.Email
		fmt.Println(e)
		if e != nil {
			fmt.Println(err)
			services.SendFront(w,models.Errors["401"],401)
			return
		}
		fmt.Println(user)
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			services.SendFront(w,models.Errors["401"],401)
			return
		}
		fmt.Println("connexion reussit")
		// Générez un token JWT ou une autre méthode pour maintenir la session

		er := SessionService.SessionStart(user, w)
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}
		// userService.UpdateSessionByID(id, sessionToken)

		json.NewEncoder(w).Encode(map[string]*models.User{"user": user})

	}
}

func Empty(user models.User) bool {
	return user.Email == "" || user.Password == "" || user.Firstname == "" ||
		user.Lastname == "" || user.DateOfBirth == "" ||
		user.Avatar == "" || user.Username == "" || user.Bio == ""
}
