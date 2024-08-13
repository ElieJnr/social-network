package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	Id         uint
	UserID   uuid.UUID
	ReceiverId uuid.UUID
	Msg        string
	CreatedAt  time.Time
}
