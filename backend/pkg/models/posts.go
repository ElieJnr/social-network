package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          int
	UserId      uuid.UUID
	Title       string
	Content     string
	Image       []byte
	CreatedAt   time.Time
	Comments    []Comment
	Like        uint
	Dislike     uint
	Statut      string
	LikeDislike []LikesDislikes
}

type Author struct {
	Firstname string
	Lastname  string
	Username  string
	Avatar    string
}

type Posts struct {
	PostID        int
	UserID        int
	Author        Author
	Content       string
	Image_url     string
	Creation_date time.Time
	Formated_date string
	Post_status   string
	Can_see       bool
	Like_nbr      int
	Dislike_nbr   int
	Comments_nbr  int
	// Comments       []Comments
	Like_status    bool
	Dislike_status bool
}
