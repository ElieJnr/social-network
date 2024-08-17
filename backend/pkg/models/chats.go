package models

import (
	"time"
)

type Chat struct {
	Id         uint      `json:"id"`
	UserID     int       `json:"user_id"`
	ReceiverId int       `json:"receiver_id"`
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
	SenderId   string
	ReceiverId string
}