package middlewares

import (
	"fmt"
	"net/http"
)

// CORSMiddleware ajoute les en-têtes CORS aux réponses HTTP
// func CORSMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Définir les en-têtes CORS
// 		w.Header().Set("Access-Control-Allow-Origin", "*") // Assurez-vous que c'est l'origine correcte
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// w.Header().Set("Access-Control-Allow-Origin", "*")
// 		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// Si c'est une requête OPTIONS (preflight), répondre immédiatement
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		// Passer à la suite (le prochain middleware ou le handler final)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func CORSMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Handling CORS for", r.Method, r.URL.Path)

// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		// w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling CORS for", r.Method, r.URL.Path)

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Répondre immédiatement aux requêtes OPTIONS
		if r.Method == http.MethodOptions {
			fmt.Println("here------------------------------------------")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
