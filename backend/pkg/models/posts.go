package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          int
	UserId      uuid.UUID
	Title       string
	Content     string
	Image       []byte
	CreatedAt   time.Time
	Comments    []Comment
	Like        uint
	Dislike     uint
	Statut		string
	LikeDislike []LikesDislikes
}
