package models

type Group struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
	CreateAt    string `json:"createAt"`
	IsMember    bool   `json:"isMember"`
}

type NewMember struct {
	UserId  string `json:"userId"`
	GroupId string `json:"groupID"`
	Status  string `json:"status"`
}
