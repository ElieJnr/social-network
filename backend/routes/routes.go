package routes

import (
	// "database/sql"
	// "net/http"
	"socialNetwork/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	// Créer un nouveau routeur
	router := mux.NewRouter()

	// Définir les routes et les associer aux handlers
	router.HandleFunc("/", handlers.HomeHandler()).Methods("GET")
	router.HandleFunc("/users", handlers.UsersHandler()).Methods("GET")
	router.HandleFunc("/signin", handlers.RegistrationHandler())
	router.HandleFunc("/login", handlers.LoginHandler())
	router.HandleFunc("/post", handlers.PostHandler()).Methods("GET")
	router.HandleFunc("/post/create", handlers.PostCreateHandler()).Methods("POST")
	router.HandleFunc("/ws", handlers.WebsocketHandler)

	// Retourner le routeur configuré
	return router
}
