package models

type Group struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string
	CreateAt    string
}

type NewMember struct {
	UserId  string `json:"userId"`
	GroupId string `json:"groupID"`
	Status string
}