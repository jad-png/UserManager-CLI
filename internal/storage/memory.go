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

func (m *Memory) GetById(id string) (*models.User, error) {
	m.um.RLock()
	defer m.um.RUnlock()

	user, ok := m.users[id]
	if !ok {
		return nil, models.ErrUserNotFound
	}
	return user, nil
}
