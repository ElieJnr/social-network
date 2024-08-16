package models

import "github.com/google/uuid"

type Member struct {
	Id      uuid.UUID
	UserID  uuid.UUID
	GroupId uuid.UUID
}
