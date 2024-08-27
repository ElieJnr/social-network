package routes

import (
	// "database/sql"
	// "net/http"

	"socialNetwork/handlers"

	"socialNetwork/middlewares"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// users
	router.Handle("/", middlewares.AuthMiddleware(handlers.HomeHandler())).Methods("GET")
	router.Handle("/users", middlewares.AuthMiddleware(handlers.UsersHandler())).Methods("GET")

	// Authentification
	router.Handle("/login", handlers.LoginHandler())
	router.Handle("/signin", handlers.RegistrationHandler())

	// posts
	router.Handle("/posts", middlewares.AuthMiddleware(handlers.PostHandler())).Methods("GET")
	router.Handle("/post/create", middlewares.AuthMiddleware(handlers.CreatePostHandler())).Methods("POST")
	
	// group


	// comments
	router.Handle("/comment/create", middlewares.AuthMiddleware(handlers.CreateCommentHandler())).Methods("POST")

	// chat
	router.HandleFunc("/ws", handlers.WebsocketHandler)


	router.Use(middlewares.CORSMiddleware)
	return router
}
