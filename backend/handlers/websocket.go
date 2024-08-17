package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"

	"github.com/gorilla/websocket"
)

// upgrader du websocket et table de map des utilsateur connecter
var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ClientWebSocketConnections = make(map[string]*websocket.Conn)
)

// service passant a travers le websocket
var (
// exemple service_Notif = services.NewNotifService() etc...
)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	//------------ Dés que l'utilisateur se connecte il est brancher au websocket via ws:localhost:port/?userId=10521@-fnc...
	// on récupére du userId de l'expéditeur ------
	UserId := r.URL.Query().Get("userId")

	// ------------------------------------------
	// ajout de l'utilisateur dans le tableau des connexions
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("probleme lors de l'initialisation: ", err)
		return
	}
	ClientWebSocketConnections[UserId] = conn

	// ------------------------------------------
	go Reader(conn)
}

func Reader(conn *websocket.Conn) error {
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			return fmt.Errorf("Json error %w", err)
		}

		switch msg.Type { // les fonction qui utiliseront la base de donnee doivent etre des services
		case "groupeChat":
			//fonction qui gere groupChat:
		case "userChat":
			//fonction qui gere userChat
		case "notifications":
			//fonction qui gere notifications
		}

	}

}

