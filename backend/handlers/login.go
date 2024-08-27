package handlers
import (
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
			EmailOrUsername string
			Password        string
		}
		credentials.EmailOrUsername = r.FormValue("emailOrUsername")
		credentials.Password = r.FormValue("password")
		user, err := UserService.UserExists(credentials.EmailOrUsername)
		// id := user.Id
		// username := user.Username
		//  := user.Password
		// email := user.Email
		if err != nil {
			fmt.Println(err)
			services.SendFront(w, models.Errors["401"], 401)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
		if err != nil {
			services.SendFront(w, models.Errors["401"], 401)
			return
		}
		fmt.Println("connexion reussit")
		// Générez un token JWT ou une autre méthode pour maintenir la session
		er := SessionService.SessionStart(user, w)
		if er != nil {
			services.SendFront(w, models.Errors["500"], 500)
			return
		}
		// userService.UpdateSessionByID(id, sessionToken)
		services.SendFront(w, map[string]*models.User{"user": user}, 200)
	}
}
func Empty(user models.User) bool {
	return user.Email == "" || user.Password == "" || user.Firstname == "" ||
		user.Lastname == "" || user.DateOfBirth == "" ||
		user.Username == "" || user.Bio == ""
}