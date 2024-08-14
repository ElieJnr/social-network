package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	Id         uint
	UserID   uuid.UUID
	ReceiverId uuid.UUID
	Msg        string
	CreatedAt  time.Time
}

var clientWebSocketConnections = make(map[int]*websocket.Conn)

func WebsocketService(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cookieName")
	if err != nil {
		return
	}
	sender := GetSender(cookie.Name)
	fmt.Println("sender name: ", sender)
}

func GetSender(cookie string) int {
	return 0
}
