package models

import "database/sql"

type Session struct {
	ID        int
	UserID    int
	Token     string //only set when creating a new session
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {

	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
