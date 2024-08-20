package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID      `json:"id"`
	Username     string         `json:"username"`
	Age          int            `json:"age"`
	Firstname    string         `json:"firstname"`
	Lastname     string         `json:"lastname"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Genre        string         `json:"genre"`
	DateOfBirth  string         `json:"dateOfBirth"`
	Bio          string         `json:"bio"`
	Avatar       string         `json:"avatar"`
	IsPrivate    bool           `json:"isPrivate"`
	Followers    []Follower     `json:"followers"`
	Session      string         `json:"session"`
	Notification []Notification `json:"notifications"`
}
