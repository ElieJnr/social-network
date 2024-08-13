package models

import "github.com/google/uuid"

type Member struct {
	Id      uint
	UserID  uuid.UUID
	GroupId uuid.UUID
}
