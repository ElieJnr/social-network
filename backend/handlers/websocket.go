package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"

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

// handler du websocket
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	//------------ Dés que l'utilisateur se connecte il est brancher au websocket via ws:localhost:port/?userId=10521@-fnc...
	// ------on récupére du userId de l'expéditeur ------
	UserId := r.URL.Query().Get("userId")
	fmt.Println("userID: ", UserId)
	// ------------------------------------------
	// ajout de l'utilisateur dans le tableau des connexions
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("probleme lors de l'initialisation: ", err)
		return
	}
	ClientWebSocketConnections[UserId] = conn
	fmt.Println("clients websocket:", ClientWebSocketConnections)

	// ------------------------------------------
	go Reader(conn)
}

func Reader(conn *websocket.Conn) error {
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			return fmt.Errorf("json error %w", err)
		}

		switch msg.Type { // les fonction qui utiliseront la base de donnee doivent etre des services
		case "groupeChat":
			//fonction qui gere groupChat:
		case "clickOnUser":
			messageService := services.NewChatService()
			err := messageService.SendStockedMessage(conn, msg.SenderId, msg.ReceiverId)
			if err != nil {
				return err
			}
		case "notifications":
			//fonction qui gere notifications
		case "sendMessage":
			messageService := services.NewChatService()
			err := messageService.RegisterMsg(msg)
			if err != nil {
				return err
			}
			sendError := conn.WriteMessage(websocket.TextMessage, []byte(msg.Content))
			if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
				sendErrorReceiver := receiverConn.WriteMessage(websocket.TextMessage, []byte(msg.Content))
				if sendErrorReceiver != nil {
					return fmt.Errorf("problem sending message to receiver: %s", sendErrorReceiver)
				}
			}
			if sendError != nil {
				return fmt.Errorf("problem sending message to users:%s", sendError)
			}
		}
	}
}
