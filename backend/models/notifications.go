package models

import "github.com/google/uuid"



//\\ A revoir //\\

type Notification struct {
	ID         uint      `json:"id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	SenderID   uuid.UUID `json:"sender_id"`
	Count      uint      `json:"count"`
}
