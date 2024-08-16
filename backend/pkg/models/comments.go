package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uuid.UUID
	PostId    uuid.UUID
	UserId    uuid.UUID
	Content   string
	CreatedAt time.Time
}
