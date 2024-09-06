package handlers

import (
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"

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

// handler du websocket
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	//------------ Dés que l'utilisateur se connecte il est brancher au websocket via ws:localhost:port/?userId=10521@-fnc...
	// ------on récupére du userId de l'expéditeur ------
	messageService := services.NewChatService()
	sender, sendErr := messageService.GetConnectedUserId(r)
	if sendErr != nil {
		return
	}
	// ------------------------------------------
	// ajout de l'utilisateur dans le tableau des connexions
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("probleme lors de l'initialisation: ", err)
		return
	}
	ClientWebSocketConnections[sender] = conn
	// fmt.Println("clients websocket:", ClientWebSocketConnections)

	// ------------------------------------------
	go Reader(conn, w, r)
}

func Reader(conn *websocket.Conn, w http.ResponseWriter, r *http.Request) error {
	sender, _ := ChatService.GetConnectedUserId(r)

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			return fmt.Errorf("json error %w", err)
		}
		// fmt.Println("message", msg)
		switch msg.Type { // les fonction qui utiliseront la base de donnee doivent etre des services

		case "groupeChat":
			//fonction qui gere groupChat:
		case "clickOnUser":
			err := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, false, false)
			if err != nil {
				fmt.Println("erreur :", err)
				return err
			}
		case "notifications":
			if msg.SubType == "read" {
				err := NotifService.MarkAsRead(msg.ReceiverId) //id of notif
				if err != nil {
					return fmt.Errorf("error read: %w", err)
				}

			}
			if msg.SubType == "sendFollow" {

				idNotif, er := utils.GenerateUuid()
				if er != nil {
					return fmt.Errorf("error read: %w", err)
				}

				notif := models.Notification{
					Id:         idNotif,
					ReceiverID: msg.ReceiverId,
					SenderID:   sender,
					Type:       "follow",
					Message:    msg.Content,
				}
				e := NotifService.CreateNotification(&notif)

				if e != nil {
					return fmt.Errorf("error read: %w", err)
				}
				fmt.Println(msg.ReceiverId)
				if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
					receiverConn.WriteJSON(msg)
				}

			}
			if msg.SubType == "addGroupe" {
				idNotif, er := utils.GenerateUuid()
				if er != nil {
					return fmt.Errorf("error read: %w", err)
				}

				author, err := utils.GetAuthor(NotifService.GetDB(), sender)

				mess := author.Firstname + " " + author.Lastname + " veut rejoindre le groupe " + msg.Content
				if err != nil {
					return fmt.Errorf("error read: %w", err)
				}

				notif := models.Notification{
					Id:         idNotif,
					ReceiverID: msg.ReceiverId,
					SenderID:   sender,
					Type:       "addGroupe",
					Message:    mess,
				}
				e := NotifService.CreateNotification(&notif)
				if e != nil {
					return fmt.Errorf("error read: %w", err)
				}
				if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
					receiverConn.WriteJSON(msg)
				}

				newMember := models.NewMember{
					UserId:  sender,
					GroupId: msg.GroupeId,
					Status:  "waiting",
				}

				e = GroupeService.AddNewMember(newMember)
				if e != nil {
					return fmt.Errorf("error on the status waiting in ws: %w", err)
				}

			}
		case "sendMessage":

			msg.SenderId = sender

			idNotif, er := utils.GenerateUuid()
			if er != nil {
				return fmt.Errorf("error read: %w", err)
			}
			notif := models.Notification{
				Id:         idNotif,
				ReceiverID: msg.ReceiverId,
				SenderID:   msg.SenderId,
				Type:       "msg",
				Message:    msg.Content,
			}
			e := NotifService.CreateNotification(&notif)

			if e != nil {
				return fmt.Errorf("error read: %w", err)
			}

			err := ChatService.RegisterMsg(msg)
			if err != nil {
				fmt.Println("error: ", err)
				return err
			}

			errToSender := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, true, false)
			if errToSender != nil {
				return errToSender
			}

			if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
				errtoreceiver := ChatService.SendStockedMessage(receiverConn, sender, msg.ReceiverId, true, false)
				if errtoreceiver != nil {
					return errtoreceiver
				}
			}

		case "userSender":
			// fmt.Println("senderId :", msg.SenderId, "sender: ", sender, "receiver: ", msg.ReceiverId)
			user, err := ChatService.FetchUser(ClientWebSocketConnections, sender)
			if err != nil {
				fmt.Println("err: ", err)
				return fmt.Errorf("error : %w", err)
			}
			conn.WriteMessage(websocket.TextMessage, []byte(user))

		case "enterGroupMessage":
			err := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, false, true)
			if err != nil {
				return err
			}
		case "newGroupChat":
			msg.SenderId = sender
			err := ChatService.RegisterMsg(msg)
			if err != nil {
				fmt.Println("error: ", err)
				return err
			}
			er := ChatService.SendStockedMessage(conn, msg.SenderId, msg.ReceiverId, false, true)
			if er != nil {
				return err
			}
		}
	}
}
