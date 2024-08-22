package models

import (
	"github.com/google/uuid"
)

type Follower struct {
	Id           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"userId"`
	FollowedUser uuid.UUID `json:"followedUser"`
}
