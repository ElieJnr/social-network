package middlewares

import (
	"context"
	"net/http"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
)

// AuthMiddleware est un middleware qui vérifie la validité du token et récupère l'utilisateur associé
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer le cookie "session_token"
		// tokenCookie, err := r.Cookie("session_token")
		// if err != nil {
		// 	http.Error(w, "Missing or invalid session token", http.StatusUnauthorized)
		// 	return
		// }

		// // Vérifier le token dans la base de données
		// user := models.User{}
		// db := sqlite.GlobalDB
		// err = db.GetDB().QueryRow("SELECT id, firstname, email FROM users WHERE id = ?", tokenCookie.Value).Scan(&user.Id, &user.Firstname, &user.Email)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }
		var sessService = services.NewSessionService()
		_, user, err := sessService.Authenticated(w, r)
		if err != nil {
			return
		}

		// Ajouter l'utilisateur au contexte de la requête
		ctx := context.WithValue(r.Context(), utils.UserContextKey, &user)

		// Créer une nouvelle requête avec le contexte modifié
		r = r.WithContext(ctx)

		// Passer à la suite (le prochain middleware ou le handler final)
		next.ServeHTTP(w, r)
	})
}
