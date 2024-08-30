package handlers

import (
	"socialNetwork/pkg/services"
)

var (
	UserService    *services.UserService
	LikeService    *services.LikeService
	PostService    *services.PostService
	ChatService    *services.ChatService
	NotifService   *services.NotifService
	SessionService *services.SessionService
	CommentService *services.CommentService
	GroupeService *services.GroupeService
	MemberService *services.MemberService
)
