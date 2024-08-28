package models

import (
	"time"
)

type Author struct {
	Firstname string
	Lastname  string
	Username  string
	Avatar    string
}

type Posts struct {
	PostID         string
	UserID         string
	Content        string
	Image_url      string
	Post_status    string
	Creation_date  time.Time
	Formated_date  string
	Author         Author
	Can_see        bool
	Like_nbr       int
	Dislike_nbr    int
	Comments_nbr   int
	Comments       []Comment
	// OwnPost        []Posts
	Like_status    bool
	Dislike_status bool
}

type CheckResult struct {
	Success      bool
	Error        string
	Content      string
	PhotoURL     string
	Status       string
	AllowedUsers []string
}
