package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/services"
)

type CookieRequest struct {
	Cookie string `json:"cookie"`
}

func ValidateCookieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request CookieRequest

		// Décoder la requête
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		fmt.Println("thecookies", request.Cookie)

		// Logique pour valider le cookie
		valid := validateCookie(request.Cookie)

		if !valid{
			services.DelCookie(w)
		}

		// Répondre avec le résultat de la validation
		json.NewEncoder(w).Encode(map[string]bool{"valid": valid})
	}
}

func validateCookie(token string) bool {
	_, err := SessionService.CheckSession(token)
	return err==nil
}
