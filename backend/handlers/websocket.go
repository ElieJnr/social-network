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
		fmt.Println("probleme lors de l'initialisation: %w", err)

	}
	ClientWebSocketConnections[sender] = conn
	// fmt.Errorf("clients websocket:", ClientWebSocketConnections)

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
		switch msg.Type { // les fonction qui utiliseront la base de donnee doivent etre des services

		case "groupeChat":
			//fonction qui gere groupChat:
		case "clickOnUser":
			err := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, false, false)
			if err != nil {
				return fmt.Errorf("erreur: %w", err)

			}
		case "notifications":
			if msg.SubType == "read" {
				err := NotifService.MarkAsRead(msg.ReceiverId) //id of notif
				if err != nil {
					return fmt.Errorf("error read: %w", err)
				}

			}
			if msg.SubType == "invitation" {
				idNotif, er := utils.GenerateUuid()
				if er != nil {
					return fmt.Errorf("error read: %w", err)
				}
				author, err := utils.GetAuthor(NotifService.GetDB(), sender)

				if err != nil {
					return fmt.Errorf("error read: %w", err)
				}
				var title string
				errr := NotifService.GetDB().QueryRow("SELECT title FROM Groups WHERE id = ?", msg.GroupeId).Scan(&title)
				if errr != nil {
					return fmt.Errorf("error read: %w", errr)
				}
				notif := models.Notification{
					Id:         idNotif,
					ReceiverID: msg.ReceiverId,
					SenderID:   sender,
					Type:       "invitation",
					Message:    author.Firstname + " " + author.Lastname + " vous invite a rejoindre le groupe:" + title,
					GroupId: msg.GroupeId,
				}
				e := NotifService.CreateNotification(&notif)

				if e != nil {
					return fmt.Errorf("error read: %w", err)
				}

				if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
					receiverConn.WriteJSON(msg)
				}
			}
			if msg.SubType == "event" {
				members, e := MemberService.GetMembership(msg.GroupeId)
				if e != nil {
					return fmt.Errorf("error member: %w", err)
				}

				for _, m := range members {
					if m.UserId != sender {
						idNotif, er := utils.GenerateUuid()
						if er != nil {
							return fmt.Errorf("error read: %w", err)
						}
						notif := models.Notification{
							Id:         idNotif,
							ReceiverID: m.UserId,
							SenderID:   sender,
							Type:       "event",
							Message:    "a créer un nouveau evenement intitulé sg : " + msg.Content,
						}
						e := NotifService.CreateNotification(&notif)

						if e != nil {
							return fmt.Errorf("error read: %w", err)
						}

						if receiverConn, ok := ClientWebSocketConnections[m.UserId]; ok {
							receiverConn.WriteJSON(msg)
						}
					}
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

				groupOwnerID, err:= MemberService.GetGroupOwnerID(msg.GroupeId)

				notif := models.Notification{
					Id:         idNotif,
					ReceiverID: groupOwnerID,
					SenderID:   sender,
					Type:       "addGroupe",
					Message:    mess,
					GroupId:    msg.GroupeId,
				}
				e := NotifService.CreateNotification(&notif)
				if e != nil {
					return fmt.Errorf("error read: %w", err)
				}
				if receiverConn, ok := ClientWebSocketConnections[groupOwnerID]; ok {
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
				return fmt.Errorf("error: %w ", err)
			}

			errToSender := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, true, false)
			if errToSender != nil {
				return fmt.Errorf("error: %w", errToSender)
			}

			if receiverConn, ok := ClientWebSocketConnections[msg.ReceiverId]; ok {
				errtoreceiver := ChatService.SendStockedMessage(receiverConn, sender, msg.ReceiverId, true, false)
				if errtoreceiver != nil {
					return fmt.Errorf("error: %w", errtoreceiver)
				}
			}

		case "userSender":
			// fmt.Errorf("senderId :", msg.SenderId, "sender: ", sender, "receiver: ", msg.ReceiverId)
			user, err := ChatService.FetchUser(ClientWebSocketConnections, sender)
			if err != nil {
				return fmt.Errorf("err: %w ", err)
			}
			conn.WriteMessage(websocket.TextMessage, []byte(user))

		case "enterGroupMessage":
			err := ChatService.SendStockedMessage(conn, sender, msg.ReceiverId, false, true)
			if err != nil {
				return fmt.Errorf("error: %w", err)
			}
		case "newGroupChat":
			msg.SenderId = sender

			member, e := ChatService.GetGroupMembership(msg.ReceiverId)
			if e != nil {
				return fmt.Errorf("error: %w", e)
			}
			// fmt.Errorf("member: ", member)

			err := ChatService.RegisterMsg(msg)
			if err != nil {
				return fmt.Errorf("error: %w", err)
			}

			for _, m := range member {
				if conn, ok := ClientWebSocketConnections[m]; ok {
					er := ChatService.SendStockedMessage(conn, msg.SenderId, msg.ReceiverId, false, true)
					if er != nil {
						return fmt.Errorf("error: %w", er)
					}
				}
			}
		}
	}
}
