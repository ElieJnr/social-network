package routes

import (
	// "database/sql"
	// "net/http"
	"socialNetwork/handlers"
	"socialNetwork/pkg/db/sqlite"

	"github.com/gorilla/mux"
)

func InitializeRoutes(db *sqlite.DB) *mux.Router {
	// Créer un nouveau routeur
	router := mux.NewRouter()

	// Définir les routes et les associer aux handlers
	router.HandleFunc("/", handlers.HomeHandler(db)).Methods("GET")
	router.HandleFunc("/users", handlers.UsersHandler(db)).Methods("GET")
	router.HandleFunc("/signin",handlers.RegistrationHandler())
	router.HandleFunc("/login",handlers.LoginHandler())
	router.HandleFunc("/post", handlers.PostHandler(db)).Methods("GET")
	router.HandleFunc("/post/create", handlers.PostCreateHandler(db)).Methods("POST")

	// Retourner le routeur configuré
	return router
}
