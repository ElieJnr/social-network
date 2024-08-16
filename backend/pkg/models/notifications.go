package models

import "github.com/google/uuid"



//\\ A revoir //\\

type Notification struct {
	ID         uuid.UUID      `json:"id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	SenderID   uuid.UUID `json:"sender_id"`
	Count      uuid.UUID      `json:"count"`
}
