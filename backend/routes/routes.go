package routes

import (
	// "database/sql"
	// "net/http"

	"fmt"
	"socialNetwork/handlers"

	"socialNetwork/middlewares"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	fmt.Println("Initializing routes")
	// Créer un nouveau routeur
	router := mux.NewRouter()

	// Définir les routes et les associer aux handlers
	router.Handle("/", middlewares.AuthMiddleware(handlers.HomeHandler())).Methods("GET")
	router.Handle("/users", middlewares.AuthMiddleware(handlers.UsersHandler())).Methods("GET")
	router.Handle("/post", middlewares.AuthMiddleware(handlers.PostHandler())).Methods("GET")
	router.Handle("/post/create", middlewares.AuthMiddleware(handlers.CreatePostHandler())).Methods("POST")
	router.Handle("/comment/create", middlewares.AuthMiddleware(handlers.CreateCommentHandler())).Methods("POST")
	router.Handle("/signin", handlers.RegistrationHandler())
	router.Handle("/login", handlers.LoginHandler())
	router.HandleFunc("/ws", handlers.WebsocketHandler)
	router.Use(middlewares.CORSMiddleware)
	// Retourner le routeur configuré
	return router
}
