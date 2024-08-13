package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID
	Username     string
	Age          int
	Firstname    string
	Lastname     string
	Email        string
	Password     string
	Genre        string
	Bio          string
	Avatar       string
	IsPrivate	 bool
	Followers	 []Follower
	Notification []Notification `json:"notifications"`
}
