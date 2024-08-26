package routes

import (
	// "database/sql"
	// "net/http"

	"socialNetwork/handlers"

	"socialNetwork/middlewares"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	// Créer un nouveau routeur
	router := mux.NewRouter()

	// Définir les routes et les associer aux handlers

	// Supposons que UsersHandler() retourne un http.Handlers
	router.Handle("/", middlewares.AuthMiddleware(handlers.HomeHandler())).Methods("GET")
	// Supposons que UsersHandler() retourne un http.Handler
	// router.Handle("/users", middlewares.AuthMiddleware()handlers.UsersHandler()).Methods("GET")
	router.Handle("/users", middlewares.AuthMiddleware(handlers.UsersHandler())).Methods("GET")
	router.Handle("/posts", middlewares.AuthMiddleware(handlers.PostHandler())).Methods("GET")
	router.Handle("/post/create", middlewares.AuthMiddleware(handlers.CreatePostHandler())).Methods("POST")
	router.Handle("/comment/create", middlewares.AuthMiddleware(handlers.CreateCommentHandler())).Methods("POST")
	router.Handle("/signin", handlers.RegistrationHandler())
	router.Handle("/login", handlers.LoginHandler())
	router.HandleFunc("/ws", handlers.WebsocketHandler)
	router.Use(middlewares.CORSMiddleware)
	// Retourner le routeur configuré
	return router
}
