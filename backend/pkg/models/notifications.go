package models

type Notification struct {
	Id         string
	ReceiverID string
	SenderID   string
	Type       string
	Message    string
	IsRead     bool
	CreateAt   string
	SenderInfo Author
}

type MessageNotif struct {
	IdNotif         string
	Desc    string
	SubType     string
}
