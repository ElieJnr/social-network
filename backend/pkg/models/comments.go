package models

import (
	"time"
)

type Comment struct {
	CommentID     string
	UserID        string
	Content       string
	Image_url     string
	Creation_date time.Time
	Formated_date string
	Author        Author
}
