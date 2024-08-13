package models

import (
	"github.com/google/uuid"
)

type Follower struct {
	Id           uint
	UserID       uuid.UUID
	FollowedUser uuid.UUID
}
