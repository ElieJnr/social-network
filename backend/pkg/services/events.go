package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"

	"github.com/gofrs/uuid/v5"
)

type EventService struct {
	db *sql.DB
}

func NewEventService() *EventService {
	dbs := sqlite.GlobalDB
	return &EventService{
		db: dbs.GetDB(),
	}
}

func (events *EventService) CreateEvent(event models.Event) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("could not create event ID: %w", err)
	}
	query := `INSERT INTO Events (id, memberId, groupId, title, description, eventDate)
               VALUES (?, ?, ?, ?, ?, ?)`
	_, err = events.db.Exec(query, id.String(), event.MemberID, event.GroupId, event.Title, event.Description, event.EventDate)
	if err != nil {
		return fmt.Errorf("could not create event: %w", err)
	}

	return nil
}

func (events *EventService) RespondToEvent(response models.EventResponse) error {
	query := `INSERT INTO EventResponses (eventId, memberId, response)
              VALUES (?, ?, ?)`
	_, err := events.db.Exec(query, response.EventID, response.MemberID, response.Response)
	if err != nil {
		return fmt.Errorf("could not respond to event: %w", err)
	}

	return nil
}

func (events *EventService) GetEventsByGroupId(groupId string) ([]models.Event, error) {
	query := `SELECT id, memberId, groupId, title, description, eventDate
              FROM Events
              WHERE groupId = ?
              ORDER BY eventDate ASC`

	rows, err := events.db.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("could not query events: %w", err)
	}
	defer rows.Close()

	var eventsList []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.MemberID, &event.GroupId, &event.Title, &event.Description, &event.EventDate)
		if err != nil {
			return nil, fmt.Errorf("could not scan event: %w", err)
		}
		eventsList = append(eventsList, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over events: %w", err)
	}
	fmt.Println("eventss", eventsList)
	return eventsList, nil
}

// GetResponseByEventAndMember retrieves a response for a specific event by a specific member
func (s *EventService) GetResponseByEventAndMember(groupId string, memberId string) ([]models.EventResponse, error) {
	query := `
		SELECT er.id, er.eventId, er.memberId, er.response
		FROM EventResponses er
		JOIN Events e ON er.eventId = e.id
		WHERE e.groupId = ? AND er.memberId = ?
	`
	rows, err := s.db.Query(query, groupId, memberId)
	if err != nil {
		return nil, fmt.Errorf("could not query responses: %w", err)
	}
	defer rows.Close()

	var responses []models.EventResponse
	for rows.Next() {
		var response models.EventResponse
		err := rows.Scan(&response.ID, &response.EventID, &response.MemberID, &response.Response)
		if err != nil {
			return nil, fmt.Errorf("could not scan response: %w", err)
		}
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over responses: %w", err)
	}

	return responses, nil
}

// UpdateResponse updates the response of a member for a specific event
func (s *EventService) UpdateResponse(response models.EventResponse) error {
	query := `UPDATE EventResponses SET response = ? WHERE eventId = ? AND memberId = ?`
	_, err := s.db.Exec(query, response.Response, response.EventID, response.MemberID)
	if err != nil {
		return fmt.Errorf("could not update response: %w", err)
	}
	return nil
}
