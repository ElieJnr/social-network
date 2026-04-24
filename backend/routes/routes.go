package routes

import (
	"net/http"
	"socialNetwork/handlers"
	"socialNetwork/middlewares"
)

func InitializeRoutes() {
	// Auth middleware
	authMiddleware := middlewares.AuthMiddleware

	// Users routes
	http.HandleFunc("/users", middlewares.CORSMiddleware((authMiddleware(handlers.UsersHandler()))))
	http.HandleFunc("/getUser", middlewares.CORSMiddleware(authMiddleware(handlers.GetUsertHandler())))
	http.HandleFunc("/update", middlewares.CORSMiddleware(authMiddleware(handlers.Edit())))

	// Follower routes
	http.HandleFunc("/follow", middlewares.CORSMiddleware(authMiddleware(handlers.Follow())))

	// Authentication routes
	http.HandleFunc("/login", middlewares.CORSMiddleware(handlers.LoginHandler()))
	http.HandleFunc("/validatecookie", middlewares.CORSMiddleware(handlers.ValidateCookieHandler()))
	http.HandleFunc("/signin", middlewares.CORSMiddleware(handlers.RegistrationHandler()))
	http.HandleFunc("/logout", middlewares.CORSMiddleware(authMiddleware(handlers.Logout())))

	// Posts routes
	http.HandleFunc("/posts", middlewares.CORSMiddleware(authMiddleware(handlers.PostHandler("allPost"))))
	http.HandleFunc("/postUser", middlewares.CORSMiddleware(authMiddleware(handlers.PostHandler("userPost"))))
	http.HandleFunc("/post/create", middlewares.CORSMiddleware(authMiddleware(handlers.CreatePostHandler())))

	// Group routes
	http.HandleFunc("/group/addNewMemberToGroup", middlewares.CORSMiddleware(handlers.AddNewMember()))
	http.HandleFunc("/group/notAddNewMemberToGroup", middlewares.CORSMiddleware(handlers.NotAddNewMember()))
	http.HandleFunc("/group/getGroups", middlewares.CORSMiddleware(authMiddleware(handlers.GetGroups())))
	http.HandleFunc("/group/posts", middlewares.CORSMiddleware(authMiddleware(handlers.GroupPostHandler())))
	http.HandleFunc("/group/createGroup", middlewares.CORSMiddleware(authMiddleware(handlers.CreateGroups())))
	http.HandleFunc("/group/post/create", middlewares.CORSMiddleware(authMiddleware(handlers.GroupCreatePostHandler())))
	http.HandleFunc("/group/createEvent", middlewares.CORSMiddleware(authMiddleware(handlers.CreateEventHandler())))
	http.HandleFunc("/group/getEvents", middlewares.CORSMiddleware(authMiddleware(handlers.GetEventsByGroupHandler())))
	http.HandleFunc("/group/respondEvent", middlewares.CORSMiddleware(authMiddleware(handlers.RespondToEventHandler())))
	http.HandleFunc("/group/GetrespondEvent", middlewares.CORSMiddleware(authMiddleware(handlers.GetResponsesByGroupAndMemberHandler())))
	http.HandleFunc("/group/suggGroup", middlewares.CORSMiddleware(authMiddleware(handlers.SuggGroup())))

	// Reactions routes
	http.HandleFunc("/like", middlewares.CORSMiddleware(authMiddleware(handlers.LikeHandler())))

	// Comments routes
	http.HandleFunc("/comments", middlewares.CORSMiddleware(authMiddleware(handlers.CommentHandler())))
	http.HandleFunc("/comment/create", middlewares.CORSMiddleware(authMiddleware(handlers.CreateCommentHandler())))

	// Notifications
	http.HandleFunc("/notifications", middlewares.CORSMiddleware(authMiddleware(handlers.NotifHandler())))

	// WebSocket for chat (this might need special handling for CORS)
	http.HandleFunc("/ws", handlers.WebsocketHandler)
}
