package models

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type message struct {
	Type     string
	SenderId   int
	ReceiverId int
	Content  string
}

// type Chat struct {
// 	Id         uint      `json:"id"`
// 	UserID     uuid.UUID `json:"user_id"`
// 	ReceiverId uuid.UUID `json:"receiver_id"`
// 	Msg        string    `json:"msg"`
// 	CreatedAt  time.Time `json:"created_at"`
// }

// type ChatMessage struct {
// 	Type      string    `json:"type"`
// 	Message   string    `json:"message"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// var clientWebSocketConnections = make(map[uuid.UUID]*websocket.Conn)
var clientWebSocketConnections = make(map[int]*websocket.Conn)

//step
// une fois que le client a clique sur l'utilisateur auxquelles ils souhaitent envoyees le message :
// 		on l'ajoute au tableau des connexions
//		on va dans la base de donnees on fetche tout les messages dejas creer entre les deux utilisateurs
// 		chaque message envoyees a est directement envoyees au destinataire via le websocket nouveau
// 		et des que l'utilisateur quitte la discussion on le supprime direct du tableau des connexions

// creation du websockets pour la gestion du chat simple
func WebsocketService(w http.ResponseWriter, r *http.Request) {
	// recuperation de l'userId de l'expediteur
	cookie, err := r.Cookie("cookieName")
	if err != nil {
		return
	}
	senderId := GetSender(cookie.Name)// le front peut gerer ca de maniere securiser
	// ------------------------------------------
	// ajout de l'utilisateur dans le tableau des connexions
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("probleme lors de l'initialisation: ", err)
		return
	}
	clientWebSocketConnections[senderId] = conn
	// ------------------------------------------

	// passons a la recuperation de l'id du destinataire
	// aucune strategie n'a encore ete defini pour la recuperation de cette derniere
	// receiverId := GetReceiver()
	// fmt.Println("receiverID:", receiverId)
	// ------------------------------------------

	// envoyons les messages stockees dans la base de donnees a l'expediteur
	SendStockedMessage(conn)
	// ------------------------------------------
	// go Reader(conn, senderId, receiverId)
	go Reader(conn)
	//a la palce mettre go Reader(conn)
}

// reader du websockets de gestion des chats simple
// func Reader(conn *websocket.Conn, sender, receiver uuid.UUID) {
// 	for {
// 		_, data, err := conn.ReadMessage()
// 		if err != nil || len(data) == 0 {
// 			return
// 		}
// 		fmt.Println("donnees", data)

//			var newMessage Chat
//			newMessage.UserID = sender
//			newMessage.ReceiverId = receiver
//			MessageType := handleMessage(data, newMessage)
//			fmt.Println("message type", MessageType)
//			RegisterData(newMessage)
//			receiverConn, ok := clientWebSocketConnections[receiver]
//			if !ok {
//				fmt.Println("Receiver not connected")
//				continue
//			}
//			if receiverConn != nil {
//				messageJSON, err := json.Marshal(newMessage)
//				if err != nil {
//					return
//				}
//				receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
//				conn.WriteMessage(websocket.TextMessage, messageJSON)
//			}
//		}
//	}
// ICI le message vient avec l'id du sender du receiver
func Reader(conn *websocket.Conn) {
	for {
		var msg message
		//lorsque l'utilisateur envoi le message via le websocket il envoit un objet lobjet est directement lu a partir de ReadJSON
		//ainsi tous les information du message sont dedans et on peut gerer la fonctionalite selon le type de msg
		err := conn.ReadJSON(&msg)
		if err != nil {
			return
		}
		if msg.Type == "message" {
			// faire quelque chose

		}
		if msg.Type == "typing" {
			// faire quelque chose

		}

		// etc..............

		//...........
	}
}

// ici devra etre implemente la logique de recuperation de l'uuid de l'expediteur
func GetSender(cookie string) int { //------- l'ID de l'tulisateur vient avec le message via next js qui peut gerer ca de manier securiser------
	// Exemple d'UUID que vous souhaitez retourner
	// id := "550e8400-e29b-41d4-a716-446655440000"

	// // Parser l'UUID à partir de la chaîne
	// senderID, err := uuid.Parse(id)
	// if err != nil {
	// 	log.Println("Erreur lors du parsing de l'UUID :", err)
	// 	return uuid.Nil // Retourner un UUID nul en cas d'erreur
	// }

	return 0
}

// ici devra etre implemente la logique de recuperation de l'uuid du destinataire
// func GetReceiver() uuid.UUID {
// 	// Exemple d'UUID que vous souhaitez retourner
// 	id := "550e8400-e29b-41d4-a716-446655440000"

// 	// Parser l'UUID à partir de la chaîne
// 	receiverID, err := uuid.Parse(id)
// 	if err != nil {
// 		log.Println("Erreur lors du parsing de l'UUID :", err)
// 		return uuid.Nil // Retourner un UUID nul en cas d'erreur
// 	}

// 	return receiverID
// }

func SendStockedMessage(conn *websocket.Conn) {

}

// ici devra etre implemente la logique d'enregistrement des messages
func RegisterData(data message) {

}

// func handleMessage(message []byte, structure Chat) string {
// 	var chatMsg ChatMessage
// 	err := json.Unmarshal(message, &chatMsg)
// 	if err != nil {
// 		return ""
// 	}
// 	structure.Msg = chatMsg.Message
// 	structure.CreatedAt = chatMsg.CreatedAt
// 	return chatMsg.Type
// }
