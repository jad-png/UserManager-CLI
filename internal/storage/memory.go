package storage

import (
	"awesomeProject/internal/models"
	"sync"
)

type Memory struct {
	users map[stringer]*models.User
	um    sync.RWMutex
}

func NewMemort() *Memory {
	return &Memory{
		users: make(map[string]*models.User),
	}
}

func (m *Memory) Create(user *models.User) error {
	m.um.Lock()
	defer m.um.Unlock()

	m.users[user.ID] = user
	return nil
}
