package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewSession(userID string, duration time.Duration) *Session {
	now := time.Now()
	return &Session{
		ID:        uuid.NewString(),
		UserID:    userID,
		Token:     uuid.NewString(),
		ExpiresAt: now.Add(duration),
		CreatedAt: now,
	}
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
