package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uuid.UUID
	MemberId    uuid.UUID
	GroupId		uuid.UUID
	Title	    string
	Option      string
	Content		string
	CreatedAt   time.Time
}
