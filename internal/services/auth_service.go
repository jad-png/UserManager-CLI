package services

import (
	"time"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
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
