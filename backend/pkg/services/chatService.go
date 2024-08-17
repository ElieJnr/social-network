package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatService struct {
	db *sql.DB
}

func NewChatService() *ChatService {
	dbs := sqlite.GlobalDB

	return &ChatService{
		db: dbs.GetDB(),
	}
}

func (c *ChatService) GetDB() *sql.DB {
	return c.db
}

func (c *ChatService) SetDB(db *sql.DB) {
	c.db = db
}

func (c *ChatService) SendMessage(msg models.Message, websocket map[string]*websocket.Conn) error {
	receiverConn, ok := websocket[msg.ReceiverId]
	if ok {
		if err := receiverConn.WriteJSON(msg); err != nil {
			return fmt.Errorf("writing error: %w", err)
		}
	}
	RegisterError := c.RegisterMsg(msg)

	if RegisterError != nil {
		return RegisterError
	}

	return nil
}

func (c *ChatService) SendStockedMessage(conn *websocket.Conn, senderId, receiverId string) error {
	messages, err := c.GetStoredMessages(senderId, receiverId)
	if err != nil {
		return err
	}

	if er := conn.WriteJSON(messages); er != nil {
		return fmt.Errorf("error writting: %w", er)
	}
	return nil
}

func (c *ChatService) GetStoredMessages(sender, receiver string) ([]models.Chat, error) {

	query := "SELECT * FROM Chats WHERE (user_id = ? AND receiver_id = ?) OR (user_id = ? AND receiver_id = ?)"
	rows, err := c.GetDB().Query(query, sender, receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve stored messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Chat
	for rows.Next() {
		var message models.Chat
		err := rows.Scan(&message.UserID, &message.ReceiverId, &message.Msg, &message.CreatedAt)
		if err != nil {
			fmt.Println("Failed to scan stored message:", err)
			continue
		}
		messages = append(messages, message)
	}

	return messages, nil
}

// ici devra etre implemente la logique d'enregistrement des messages
func (c *ChatService) RegisterMsg(msg models.Message) error {

	idMsg := uuid.NewString()

	query := "INSERT INTO Chats (id,senderId,receverId, content) VALUES (?, ?, ?, ?)"

	_, err := c.GetDB().Exec(query, idMsg, msg.SenderId, msg.ReceiverId, msg.Content)
	if err != nil {
		return fmt.Errorf("failed to register data: %w", err)
	}
	return nil
}
