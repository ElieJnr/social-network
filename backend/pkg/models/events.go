package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uint
	Title	    string
	MemberId    uuid.UUID
	GroupId		uuid.UUID
	Option      string
	Content		string
	CreatedAt   time.Time
}
