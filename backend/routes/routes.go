package routes

import (
	"socialNetwork/handlers"
	"socialNetwork/middlewares"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// users
	router.Handle("/", middlewares.AuthMiddleware(handlers.HomeHandler())).Methods("GET")
	router.Handle("/users", middlewares.AuthMiddleware(handlers.UsersHandler())).Methods("GET")
	router.Handle("/getUser", middlewares.AuthMiddleware(handlers.GetUsertHandler())).Methods("GET")

	//follower
	router.Handle("/follow", middlewares.AuthMiddleware(handlers.Follow())).Methods("POST")

	// Authentification
	router.Handle("/login", handlers.LoginHandler()).Methods("POST")
	router.Handle("/signin", handlers.RegistrationHandler()).Methods("POST")
	router.Handle("/logout", middlewares.AuthMiddleware(handlers.Logout())).Methods("POST")
	router.Handle("/validatecookie", handlers.ValidateCookieHandler())

	// posts
	router.Handle("/posts", middlewares.AuthMiddleware(handlers.PostHandler("allPost"))).Methods("GET")
	router.Handle("/postUser", middlewares.AuthMiddleware(handlers.PostHandler("userPost"))).Methods("GET")
	router.Handle("/post/create", middlewares.AuthMiddleware(handlers.CreatePostHandler())).Methods("POST")

	// group
	router.Handle("group/posts", middlewares.AuthMiddleware(handlers.PostHandler("allPost"))).Methods("GET")
	router.Handle("group/post/create", middlewares.AuthMiddleware(handlers.CreatePostHandler())).Methods("POST")
	router.Handle("/group/createGroup", middlewares.AuthMiddleware(handlers.CreateGroups())).Methods("POST")
	router.Handle("/group/addNewMemberToGroup", handlers.AddNewMember()).Methods("POST")
	router.Handle("/group/getGroups", handlers.GetGroups()).Methods("GET")

	// Reactions
	router.Handle("/like", middlewares.AuthMiddleware(handlers.LikeHandler())).Methods("POST")

	// comments
	router.Handle("/comments", middlewares.AuthMiddleware(handlers.CommentHandler())).Methods("GET")
	router.Handle("/comment/create", middlewares.AuthMiddleware(handlers.CreateCommentHandler())).Methods("POST")

	// chat
	router.HandleFunc("/ws", handlers.WebsocketHandler)

	//notifications
	router.Handle("/notifications", middlewares.AuthMiddleware(handlers.NotifHandler())).Methods("GET")

	router.Use(middlewares.CORSMiddleware)
	return router
}
