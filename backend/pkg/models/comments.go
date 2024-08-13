package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uint
	PostId    uint
	UserId    uuid.UUID
	Content   string
	CreatedAt time.Time
}
