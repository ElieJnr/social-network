package config

import (
	"socialNetwork/handlers"
	"socialNetwork/pkg/services"
)

func InitServices() {
	handlers.PostService = services.NewPostService()
	handlers.LikeService = services.NewLikeService()
	handlers.ChatService = services.NewChatService()
	handlers.NotifService = services.NewNotifService()
	handlers.UserService = services.NewUserService()
	handlers.SessionService = services.NewSessionService()
	handlers.CommentService = services.NewCommentService()
	handlers.GroupeService = services.NewGroupeService()
	handlers.MemberService = services.NewMemberService()
}
