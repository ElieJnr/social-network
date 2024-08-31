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
)

type SessionService struct {
	db *sql.DB
}

func NewSessionService() *SessionService {
	dbs := sqlite.GlobalDB

	return &SessionService{
		db: dbs.GetDB(),
	}
}

func (s *SessionService) GetDB() *sql.DB {
	return s.db
}

func (s *SessionService) SetDB(db *sql.DB) {
	s.db = db
}

func (s *SessionService) SessionStart(user *models.User, w http.ResponseWriter) error {
	fmt.Println("session start")
    SessionToken, err := utils.GenerateUuid()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    ExpiresAT := time.Now().UTC().Add(time.Hour*24)
    expired_at := ExpiresAT.Format("2006-01-02 15:04:05")
    maxAge := 4400
    if saveErr := s.SaveSession(SessionToken, user.Username, expired_at, user.Id); saveErr != nil {
        return saveErr
    }
	fmt.Println("cookis start")
    cookie := http.Cookie{
        Name:     "session_token",
        Value:    SessionToken,
        Expires:  ExpiresAT,
        MaxAge:   maxAge,
    }
    http.SetCookie(w, &cookie)
    return nil
}


func (s *SessionService) SaveSession(token, username, expired_AT string, user_id string) error {
	_, err := s.GetDB().Exec("INSERT INTO sessions (sessionId,userId,username,expired_at) VALUES(?,?,?,?)", token, user_id, username, expired_AT)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func (s *SessionService) CheckSession(token string) (*models.Session, error) {
	row := s.GetDB().QueryRow("SELECT * FROM sessions WHERE sessionId = ?", token)
	var sess models.Session
	err := row.Scan(&sess.Token, &sess.UserId, &sess.Username, &sess.Expired_at, &sess.Create_at)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// update the session time each time the user navigates to the page
func (s *SessionService) RefreshSession(token string) error {
	ExpiresAT := time.Now().UTC().Add(time.Hour)
	expired_at := ExpiresAT.Format("2006-01-02 15:04:05")
	_, err := s.GetDB().Exec("UPDATE sessions SET expired_at = ? WHERE sessionId = ? ", expired_at, token)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func (s *SessionService) DeleteSession(tokenOrUserId string) error {
	_, err := s.GetDB().Exec("DELETE  FROM sessions WHERE sessionId = ? OR userId = ?", tokenOrUserId,tokenOrUserId)
	if err != nil {
		return fmt.Errorf("error deleting: %w", err)
	}
	return nil
}

func (s *SessionService) GetExpired_AT(SessionID string) (string, error) {
	var expired string
	row := s.GetDB().QueryRow("SELECT expired_at FROM sessions WHERE sessionId = ?", SessionID)
	err := row.Scan(&expired)
	if err != nil {
		return "", err
	}
	return expired, nil
}

func (s *SessionService) Authenticated(w http.ResponseWriter, r *http.Request) (*http.Cookie, *models.Session, error) {
	c, err := r.Cookie("session_token")
	// fmt.Println(r.Cookies())
	if err != nil {
		data := models.Data{
			NoAuth: true,
		}
		SendFront(w, data, 200)
		return nil, nil, err
	}
	sessionID := c.Value
	UserInfo, er := s.CheckSession(sessionID)
	if er != sql.ErrNoRows && er != nil {
		SendFront(w, models.Errors["500"], 500)
		return nil, nil, er
	}
	if er == sql.ErrNoRows {
		DelCookie(w)
		data := models.Data{
			NoAuth: true,
		}
		SendFront(w, data, 200)
		return nil, nil, er
	}
	//// FROM HERE THE SESSION EXISTS, LET'S CHECK IF IT IS EXPIRED OR NOT BY RECOVERING
	// EXPIRED_AT AND COMPARE IT WITH TIME.NOW()
	exp, err := s.GetExpired_AT(sessionID)
	if err != nil {
		// Error500(w)
		return nil, nil, err
	}
	expiry, err := ParseTime(exp)
	if err != nil {
		SendFront(w, models.Errors["500"], 500)
		return nil, nil, err
	}
	if expiry.Before(time.Now().UTC()) {
		// IamDisConnected(UserInfo.Username)
		e := s.DeleteSession(sessionID)
		if e != nil {
			// Error500(w)
			return nil, nil, e
		}
		DelCookie(w)
		data := models.Data{
			NoAuth: true,
		}
		fmt.Println("expired")
		SendFront(w, data, 200)
		return nil, nil, nil
	}
	return c, UserInfo, nil
}

func ParseTime(timeStr string) (*time.Time, error) {
	layout2 := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout2, timeStr)
	parsedTime = parsedTime.UTC()
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

func DelCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func SendFront(w http.ResponseWriter, api interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(api)
	if err != nil {
		// Utiliser http.Error pour envoyer un message d'erreur et un code de statut
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

// if v := cookie.String(); v != "" {
// 	w.Header().Add("Set-Cookie", v)
// }