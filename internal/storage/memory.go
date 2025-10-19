package storage

import (
	"awesomeProject/internal/models"
	"sync"
)

type MemoryStorage struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make(map[string]*models.User),
	}
}
