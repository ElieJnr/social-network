package models

import (
	"time"
)

type Author struct {
	Firstname string
	Lastname  string
	Username  string
	Avatar    string
	Email     string
	IsPrivate bool
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
	HasImage       bool
	Can_see        bool
	Like_nbr       int
	Comments_nbr   int
	IsFollower     bool
	Like_status    bool
	Dislike_status bool
}

type CheckResult struct {
	Success      bool
	Error        string
	Content      string
	PhotoURL     string
	Status       string
	PostId       string
	GroupId      string
	AllowedUsers []string
}
