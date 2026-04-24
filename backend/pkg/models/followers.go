package models

import "github.com/gofrs/uuid/v5"

type Follower struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	FollowedId uuid.UUID `json:"followedUser"`
	Statut     bool      `json:"statut"`
}
