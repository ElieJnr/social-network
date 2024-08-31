package models

import (
	"time"
)

type Chat struct {
	Id         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ReceiverId string    `json:"receiver_id"`
	Msg        string    `json:"msg"`
	CreatedAt  time.Time `json:"created_at"`
}
type ChatMessage struct {
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// ce type sera utiliser pour tous les message envoyer sur le websocket, il peut etre optimiser selon les besoins
type Message struct {
	Type       string
	Content    string
	ReceiverId string
	SenderId   string
	SubType    string
}
