package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
	"time"

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

// les messages une fois recupere sont envoyes a l'utilisateur via sa connection websocket
func (c *ChatService) SendStockedMessage(conn *websocket.Conn, senderId, receiverId string, newMessage bool) error {
	booleen := false
	if newMessage {
		booleen = true
	}
	messages, err := c.GetStoredMessages(senderId, receiverId, booleen)
	if err != nil {
		return err
	}

	if er := conn.WriteJSON(messages); er != nil {
		return fmt.Errorf("error writting: %w", er)
	}
	return nil
}

// fonction de recuperation des messages enregistre dans la base de donnee
type sendMessage struct {
	Type    string
	Message []models.Chat
}

func (c *ChatService) GetStoredMessages(sender, receiver string, newMessage bool) (sendMessage, error) {
	var sendMessage sendMessage

	query := "SELECT id, senderId, receverId, content, sendAt FROM Chats WHERE (senderId = ? AND receverId = ?) OR (receverId = ? AND senderId = ?)"
	rows, err := c.GetDB().Query(query, sender, receiver, sender, receiver)
	if err != nil {
		return sendMessage, fmt.Errorf("failed to retrieve stored messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Chat
	for rows.Next() {
		var message models.Chat
		err := rows.Scan(&message.Id, &message.UserID, &message.ReceiverId, &message.Msg, &message.CreatedAt)
		if err != nil {
			fmt.Println("Failed to scan stored message:", err)
			continue
		}
		messages = append(messages, message)
	}

	fmt.Println("messages: ", messages)
	if newMessage {
		sendMessage.Type = "sendMessage"
	}
	sendMessage.Type = "clickOnUser"
	sendMessage.Message = messages

	return sendMessage, nil
}

// fonction d'enregistrement des messages dans la base de donnees
func (c *ChatService) RegisterMsg(msg models.Message) error {

	idMsg, er := utils.GenerateUuid()
	if er != nil {
		return er
	}
	query := "INSERT INTO Chats (id,senderId,receverId, content,type) VALUES (?, ?, ?, ?,?)"

	_, err := c.GetDB().Exec(query, idMsg, msg.SenderId, msg.ReceiverId, msg.Content, "userText")
	if err != nil {
		return fmt.Errorf("failed to register data: %w", err)
	}
	return nil
}

func (c *ChatService) FetchUser(connTab map[string]*websocket.Conn, actualuser string) ([]byte, error) {
	query := `
		SELECT DISTINCT u.id, u.firstname, u.lastname 
		FROM Users u
		JOIN Followers f 
		ON (u.id = f.followedId OR u.id = f.userId)
		WHERE (f.userId = ? OR f.followedId = ?) AND u.id != ?
	`
	rows, err := c.GetDB().Query(query, actualuser, actualuser, actualuser)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	type User struct {
		Id              string
		Firstname       string
		Lastname        string
		Lastmessage     string
		LastmessageHour string
		Online          string
	}
	type sendUser struct {
		Type  string
		Users []User
	}

	var users []User
	for rows.Next() {
		var id, firstname, lastname string
		err := rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			fmt.Println("Failed to scan user:", err)
			continue
		}

		// Vérifier si l'utilisateur est connecté
		onlineStatus := "no"
		if _, ok := connTab[id]; ok {
			onlineStatus = "yes"
		}

		query := "SELECT content, sendAt FROM Chats WHERE (senderId = ? AND receverId = ?) OR (receverId = ? AND senderId = ?) ORDER BY sendAt DESC LIMIT 1"
		lastMessageRow := c.GetDB().QueryRow(query, actualuser, id, actualuser, id)

		var lastMessage string
		var lastMessageHour time.Time
		err = lastMessageRow.Scan(&lastMessage, &lastMessageHour)
		if err != nil {
			if err == sql.ErrNoRows {
				// lastMessage = fmt.Sprintf("No message between %s and %s", actualuser, id)
			} else {
				fmt.Println("Failed to retrieve last message:", err)
			}
		}

		user := User{
			Id:              id,
			Firstname:       firstname,
			Lastname:        lastname,
			Lastmessage:     lastMessage,
			LastmessageHour: utils.FormatTimeAgo(lastMessageHour),
			Online:          onlineStatus, // Assigner le statut en ligne
		}
		users = append(users, user)
	}

	sendUsers := sendUser{
		Type:  "sendUser",
		Users: users,
	}

	jsonUsers, err := json.Marshal(sendUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal users: %w", err)
	}

	return jsonUsers, nil
}

func GetCookie(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *ChatService) GetConnectedUserId(r *http.Request) (string, error) {
	var userId string
	cookie := GetCookie(r)
	query := `SELECT userId FROM sessions WHERE sessionId = ? AND expired_at > ?`

	err := c.GetDB().QueryRow(query, cookie, time.Now()).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return userId, nil
}
