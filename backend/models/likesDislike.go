package models

import "github.com/google/uuid"

type LikesDislikes struct {
	Id     uint
	PostId uint
	UserID uuid.UUID
	Like   bool
}
