package services

import (
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

func (s *AuthService) Login(user *models.User) (*models.Session, error) {
	// step 1: find user by email
	user, err := s.userStorage.GetByEmail(user.Email)
	if err != nil {
		return nil, models.ErrUserNotFound
	}

	// step 2: check password
	if err := user.CheckPassword(user.GetPasswordHash()); err != nil {
		return nil, models.ErrUserNotFound
	}

	// step 3; create seession if user was found
	session := models.NewSession(user.ID, 24*time.Hour)
	if err := s.sessionStorage.CreateSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *AuthService) Logout(token string) error {
	return s.sessionStorage.DeleteSession(token)
}
