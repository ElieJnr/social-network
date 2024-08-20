package models

import "github.com/google/uuid"

type LikesDislikes struct {
	Id     uuid.UUID
	PostId uuid.UUID
	UserID uuid.UUID
	Like   bool
}
