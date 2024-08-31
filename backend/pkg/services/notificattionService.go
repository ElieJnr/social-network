package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type NotifService struct {
	db *sql.DB
}

func NewNotifService() *NotifService {
	dbs := sqlite.GlobalDB

	return &NotifService{
		db: dbs.GetDB(),
	}
}

func (n *NotifService) GetDB() *sql.DB {
	return n.db
}

func (n *NotifService) SetDB(db *sql.DB) {
	n.db = db
}

func (n *NotifService) CreateNotification(notif *models.Notification) error {

	query := `INSERT INTO Notifications (id,user_receiver_id,user_id, type,message)
		VALUES (?,?,?,?,?)`

	_, err := n.GetDB().Exec(query, notif.Id, notif.ReceiverID, notif.SenderID, notif.Type, notif.Message)

	if err != nil {
		return fmt.Errorf("failed to create notification: %w", err)
	}
	return nil
}

// Au cas ou une notification sera lu
func (n *NotifService) MarkAsRead(idNotif string) error {

	query := `UPDATE Notifications SET is_read = 1 WHERE id = ? `

	_, err := n.GetDB().Exec(query, idNotif)
	if err != nil {
		return fmt.Errorf("failed to update notification: %w", err)
	}
	return nil
}

// Recupere l'ensemble des notifications de l'utilisateur
func (n *NotifService) GetAllNotifications(ReceiverId string) ([]models.Notification, error) {

	query := `
	SELECT * FROM Notifications 
	WHERE (user_receiver_id = ?) ORDER BY created_at DESC
	`
	rows, err := n.GetDB().Query(query, ReceiverId)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %w", err)
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		var notification models.Notification
		if err = rows.Scan(&notification.Id, &notification.ReceiverID, &notification.SenderID, &notification.Type, &notification.Message, &notification.IsRead, &notification.CreateAt); err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		senderInfo,err := utils.GetAuthor(n.db,notification.SenderID)
		if err != nil{
			return nil, fmt.Errorf("failed to get senderInfo: %w", err)
		}
		notification.SenderInfo = senderInfo
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

// Permet d'obtenir le nombre de notification de l'utilisateur
func (n *NotifService) GetNbrOfNotif(ReceiverId string) (int, error) {
	var (
		count int
		query = `SELECT COUNT(*) FROM Notifications WHERE user_receiver_id = ? AND NOT is_read`
	)

	err := n.GetDB().QueryRow(query, ReceiverId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error :%w", err)
	}
	return count, nil
}
