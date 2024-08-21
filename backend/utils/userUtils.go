package utils

import (
	"net/http"
	"socialNetwork/pkg/models"
)

type contextKey string

const UserContextKey contextKey = "user"

func CurrentUser(w http.ResponseWriter, r *http.Request) *models.Session {
	user, ok := r.Context().Value(UserContextKey).(*models.Session)
	if !ok || user == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return nil
	}
	return user
}
