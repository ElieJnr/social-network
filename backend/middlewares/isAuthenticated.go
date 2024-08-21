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
