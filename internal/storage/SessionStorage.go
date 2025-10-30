package storage

import (
	"awesomeProject/internal/models"
	"errors"
	"sync"
)

type SessionStorage struct {
	sessions map[string]*models.Session
	mu       sync.RWMutex
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		sessions: make(map[string]*models.Session),
	}
}

func (s *SessionStorage) CreateSession(session *models.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.ID] = session
	return nil
}

func (s *SessionStorage) GetSession(token string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[token]
	if !ok {
		return nil, errors.New("session not found")
	}
	return session, nil
}

func (s *SessionStorage) DeleteSession(token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.sessions[token]; !ok {
		return nil
	}

	delete(s.sessions, token)
	return nil
}

func (s *SessionStorage) DeleteExpiredSession() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for token, session := range s.sessions {
		if session.IsExpired() {
			delete(s.sessions, token)
		}
	}
}
