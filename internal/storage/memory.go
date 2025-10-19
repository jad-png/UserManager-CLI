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

func (m *Memory) GetByEmal(email string) (*models.User, error) {
	m.um.RLock()
	defer m.um.RUnlock()

	user, ok := m.users[email]
	if !ok {
		return nil, models.ErrUserNotFound
	}
	return user, nil
}

func (m *Memory) GetAll() ([]*models.User, error) {
	m.um.RLock()
	defer m.um.RUnlock()

	users := make([]*models.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}
