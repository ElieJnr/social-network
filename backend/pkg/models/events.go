package models

type Event struct {
	ID          string `json:"id"`
	MemberID    string `json:"member_id"`
	GroupId     string `json:"groupid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"date"`
}
type EventResponse struct {
	ID       int    `json:"id"`
	EventID  string `json:"eventid"`
	MemberID string `json:"member_id"`
	Response string `json:"response"`
}
