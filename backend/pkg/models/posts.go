package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Title       string
	Content     string
	Image       string
	CreatedAt   time.Time
	Comments    []Comment
	Like        int
	Dislike     int
	Statut		string
	LikeDislike []LikesDislikes
}
