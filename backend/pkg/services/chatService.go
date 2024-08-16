package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"

	"github.com/gorilla/websocket"
)

var ClientWebSocketConnections = make(map[int]*websocket.Conn)

// à remplacer par la vraie base de donnée
var dataBase *sql.DB

func WebsocketService(w http.ResponseWriter, r *http.Request) {
	// récupération de l'userId de l'expéditeur
	cookie, err := r.Cookie("cookieName")
	if err != nil {
		return
	}
	senderId := GetSender(cookie.Name)
	// ------------------------------------------
	// ajout de l'utilisateur dans le tableau des connexions
	conn, err := models.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("probleme lors de l'initialisation: ", err)
		return
	}
	ClientWebSocketConnections[senderId] = conn
	// ------------------------------------------
	// passons a la recuperation de l'id du destinataire
	// aucune strategie n'a encore ete defini pour la recuperation de cette derniere
	receiverId := GetReceiver()
	// ------------------------------------------
	// envoyons les messages stockees dans la base de donnees a l'expediteur
	SendStockedMessage(conn, senderId, receiverId)
	// ------------------------------------------
	go Reader(conn, senderId, receiverId)
}

// reader du websockets de gestion des chats simples de groupes et des notifications
func Reader(conn *websocket.Conn, sender, receiver int) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil || len(data) == 0 {
			return
		}

		var newMessage models.Chat
		MessageType := handleMessage(data, newMessage)

		if MessageType == "groupeChat" {

		} else if MessageType == "userChat" {
			Message(newMessage, sender, receiver, conn)
		} else if MessageType == "notifications" {

		}
	}
}

func Message(newMessage models.Chat, sender, receiver int, conn *websocket.Conn) {
	newMessage.UserID = sender
	newMessage.ReceiverId = receiver
	RegisterError := RegisterData(newMessage, dataBase)

	if RegisterError != nil {
		return
	}

	receiverConn, ok := ClientWebSocketConnections[receiver]
	if ok {
		messageJSON, err := json.Marshal(newMessage)
		if err != nil {
			return
		}
		receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
		conn.WriteMessage(websocket.TextMessage, messageJSON)
	}
}

// ici devra etre implemente la logique de recuperation de l'uuid de l'expediteur
func GetSender(cookie string) int {
	return 0
}

// ici devra etre implemente la logique de recuperation de l'uuid du destinataire
func GetReceiver() int {
	return 1
}
func SendStockedMessage(conn *websocket.Conn, sender, receiver int) {
	messages := GetStoredMessages(sender, receiver)
	for _, message := range messages {
		messageJSON, err := json.Marshal(message)
		if err != nil {
			continue
		}
		conn.WriteMessage(websocket.TextMessage, messageJSON)
	}
}

func GetStoredMessages(sender, receiver int) []models.Chat {
	var messages []models.Chat
	query := "SELECT * FROM Chats WHERE (user_id = ? AND receiver_id = ?) OR (user_id = ? AND receiver_id = ?)"
	rows, err := dataBase.Query(query, sender, receiver)
	if err != nil {
		fmt.Println("Failed to retrieve stored messages:", err)
		return messages
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Chat
		err := rows.Scan(&message.UserID, &message.ReceiverId, &message.Msg, &message.CreatedAt)
		if err != nil {
			fmt.Println("Failed to scan stored message:", err)
			continue
		}
		messages = append(messages, message)
	}

	return messages
}

// ici devra etre implemente la logique d'enregistrement des messages
func RegisterData(data models.Chat, dataBase *sql.DB) error {
	query := "INSERT INTO Chats (user_id, receiver_id, msg, created_at) VALUES (?, ?, ?, ?)"
	_, err := dataBase.Exec(query, data.UserID, data.ReceiverId, data.Msg, data.CreatedAt)
	if err != nil {
		fmt.Println("Failed to register data:", err)
		return err
	}
	return nil
}

func handleMessage(message []byte, structure models.Chat) string {
	var chatMsg models.ChatMessage
	err := json.Unmarshal(message, &chatMsg)
	if err != nil {
		return ""
	}
	structure.Msg = chatMsg.Message
	structure.CreatedAt = chatMsg.CreatedAt
	return chatMsg.Type
}
