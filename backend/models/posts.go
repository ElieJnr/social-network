package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          uint
	UserId      uuid.UUID
	Title       string
	Content     string
	Image       []byte
	CreatedAt   time.Time
	Comments    []Comment
	Like        uint
	Dislike     uint
	LikeDislike []LikesDislikes
}
