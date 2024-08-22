package config

import (
	"socialNetwork/handlers"
	"socialNetwork/pkg/services"
)
// fonction pour initialiser une service apres l'avoir declarer dans serviceHnadler.go de Handler
// Donc si vous ajouter un service ajouter n'oublier pas de le declarer dans serviceHandler.go
func InitServices() {
	handlers.PostService = services.NewPostService()
	handlers.ChatService  = services.NewChatService()
	handlers.NotifService = services.NewNotifService()
	handlers.UserService = services.NewUserService()
	handlers.SessionService = services.NewSessionService()
	handlers.CommentService = services.NewCommentService()
	handlers.GroupeService = services.NewGroupeService()
	handlers.MemberService = services.NewMemberService()
}
