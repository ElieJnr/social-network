package models

import (
	"github.com/google/uuid"
)

type Follower struct {
	Id           uuid.UUID
	UserID       uuid.UUID
	FollowedUser uuid.UUID
}
