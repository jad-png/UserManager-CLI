package services

import (
	"awesomeProject/internal/auth"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"time"
)

type AuthService struct {
	userStorage    storage.Storage
	sessionStorage *storage.SessionStorage
}

func NewAuthService(userStorage storage.Storage) *AuthService {
	return &AuthService{
		userStorage:    userStorage,
		sessionStorage: storage.NewSessionStorage(),
	}
}

func (s *AuthService) Login(user *models.User) (*auth.Session, error) {
	user, err := s.userStorage.GetByEmail(user.Email)
	if err != nil {
		return nil, models.ErrUserNotFound
	}

	session := auth.NewSession(user.ID, 24*time.Hour)
	if err := s.sessionStorage.CreateSession(session); err != nil {
		return nil, err
	}

	return session, nil
}
