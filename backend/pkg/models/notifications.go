package models

import "time"

type Notification struct {
	Id         string
	ReceiverID string
	SenderID   string
	Type       string
	Message    string
	IsRead     bool
	CreateAt   time.Time
	Formated_date string
	SenderInfo Author
	GroupId string
}

type MessageNotif struct {
	IdNotif         string
	Desc    string
	SubType     string
}
