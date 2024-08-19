package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type SessionModel struct {
	DB *sql.DB
}

func (s *SessionModel) AddToken(id int, token string) error {
	stmt := `UPDATE users SET token = ?, expiry = DATETIME('now', '+12 hours')
	WHERE ? = userID`

	_, err := s.DB.Exec(stmt, token, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionModel) GetToken(token string) (string, error) {
	var result bool
	var session string
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE token = ?)`

	err := s.DB.QueryRow(stmt, token).Scan(&result)
	if err != nil {
		return "", err
	}
	if result {
		stmt2 := `SELECT token FROM users WHERE token = ?`
		err = s.DB.QueryRow(stmt2, token).Scan(&session)
		if err != nil {
			return "", err
		}
	}
	return session, nil
}

func (s *SessionModel) GetNameByToken(token string) (string, error) {
	var result bool
	var firstname string
	var lastname string
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE token = ?)`

	err := s.DB.QueryRow(stmt, token).Scan(&result)
	if err != nil {
		return "", err
	}
	if result {
		stmt2 := `SELECT firstname, lastname FROM users WHERE token = ?`
		err = s.DB.QueryRow(stmt2, token).Scan(&firstname, &lastname)
		if err != nil {
			return "", err
		}
	}
	return firstname + " " + lastname, nil
}

func (s *SessionModel) RemoveToken(token string) error {
	stmt := `UPDATE users SET token = NULL, expiry = NULL WHERE token = ?`
	_, err := s.DB.Exec(stmt, token)
	return err
}

func (s *SessionModel) IsExpired(token string) bool {
	var result *time.Time
	stmt := `SELECT expiry FROM users WHERE token = ?`

	_ = s.DB.QueryRow(stmt, token).Scan(&result)

	return result.Before(time.Now())
}

func IsLogged(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	var data bool

	if err != nil || cookie.Value == "" {
		data = false
	} else {
		data = true
	}
	return data
}

func GetUUIDToken() string {
	token, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return token.String()
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
		// Raw:      "",
	}
	http.SetCookie(w, cookie)
}