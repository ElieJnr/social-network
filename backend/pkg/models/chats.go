package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Chat struct {
	Id         uint      `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ReceiverId uuid.UUID `json:"receiver_id"`
	Msg        string    `json:"msg"`
	CreatedAt  time.Time `json:"created_at"`
}

type ChatMessage struct {
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clientWebSocketConnections = make(map[uuid.UUID]*websocket.Conn)

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
	senderId := GetSender(cookie.Name)
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
	receiverId := GetReceiver()
	fmt.Println("receiverID:", receiverId)
	// ------------------------------------------

	// envoyons les messages stockees dans la base de donnees a l'expediteur
	SendStockedMessage(conn)
	// ------------------------------------------
	go Reader(conn, senderId, receiverId)
}

// reader du websockets de gestion des chats simples
func Reader(conn *websocket.Conn, sender, receiver uuid.UUID) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil || len(data) == 0 {
			return
		}
		fmt.Println("donnees", data)

		var newMessage Chat
		newMessage.UserID = sender
		newMessage.ReceiverId = receiver
		MessageType := handleMessage(data, newMessage)
		fmt.Println("message type", MessageType)
		RegisterData(newMessage)
		receiverConn, ok := clientWebSocketConnections[receiver]
		if !ok {
			fmt.Println("Receiver not connected")
			continue
		}
		if receiverConn != nil {
			messageJSON, err := json.Marshal(newMessage)
			if err != nil {
				return
			}
			receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
			conn.WriteMessage(websocket.TextMessage, messageJSON)
		}
	}
}

// ici devra etre implemente la logique de recuperation de l'uuid de l'expediteur
func GetSender(cookie string) uuid.UUID {
	// Exemple d'UUID que vous souhaitez retourner
	id := "550e8400-e29b-41d4-a716-446655440000"

	// Parser l'UUID à partir de la chaîne
	senderID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Erreur lors du parsing de l'UUID :", err)
		return uuid.Nil // Retourner un UUID nul en cas d'erreur
	}

	return senderID
}

// ici devra etre implemente la logique de recuperation de l'uuid du destinataire
func GetReceiver() uuid.UUID {
	// Exemple d'UUID que vous souhaitez retourner
	id := "550e8400-e29b-41d4-a716-446655440000"

	// Parser l'UUID à partir de la chaîne
	receiverID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Erreur lors du parsing de l'UUID :", err)
		return uuid.Nil // Retourner un UUID nul en cas d'erreur
	}

	return receiverID
}

func SendStockedMessage(conn *websocket.Conn) {
	
}

// ici devra etre implemente la logique d'enregistrement des messages
func RegisterData(data Chat) {

}

func handleMessage(message []byte, structure Chat) string {
	var chatMsg ChatMessage
	err := json.Unmarshal(message, &chatMsg)
	if err != nil {
		return ""
	}
	structure.Msg = chatMsg.Message
	structure.CreatedAt = chatMsg.CreatedAt
	return chatMsg.Type
}
