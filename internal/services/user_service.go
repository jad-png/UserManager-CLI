package services

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
)

type Service struct {
	store storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		store: storage,
	}
}
