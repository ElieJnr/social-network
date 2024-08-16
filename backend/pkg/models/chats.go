package models

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
