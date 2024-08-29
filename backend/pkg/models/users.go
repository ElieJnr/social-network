package models

type User struct {
	Id          string     `json:"id"`
	Username    string     `json:"username"`
	Firstname   string     `json:"firstname"`
	Lastname    string     `json:"lastname"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Genre       string     `json:"genre"`
	DateOfBirth string     `json:"dateOfBirth"`
	Bio         string     `json:"bio"`
	Avatar      string     `json:"avatar"`
	IsPrivate   bool       `json:"isPrivate"`
	Followers   []Follower `json:"followers"`
	Follows     []Follower `json:"follows"`
	Notification []Notification `json:"notifications"`
}
