package storage

import (
	"awesomeProject/internal/models"
	"fmt"
	"sync"
	"time"
)

type Memory struct {
	users map[string]*models.User
	um    sync.RWMutex
}

func NewMemort() *Memory {
	return &Memory{
		users: make(map[string]*models.User),
	}
}

func (m *Memory) Create(user *models.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	if err := user.Validate(); err != nil {
		return err
	}

	m.um.Lock()
	defer m.um.Unlock()

	for _, u := range m.users {
		if u.Email == user.Email {
			return models.ErrUserExists
		}
	}

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

func (m *Memory) GetByEmail(email string) (*models.User, error) {
	m.um.RLock()
	defer m.um.RUnlock()

	for _, user := range m.users {
		if user.Email == email {
			return copyUser(user), nil
		}
	}
	return nil, models.ErrUserNotFound
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

func (m *Memory) Update(id string, user *models.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}
	m.um.Lock()
	defer m.um.Unlock()

	user, ok := m.users[id]
	if !ok {
		return models.ErrUserNotFound
	}

	user.Update(user.Name, user.Email, user.Age)

	if err := user.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *Memory) Delete(id string) error {
	m.um.Lock()
	defer m.um.Unlock()

	if _, ok := m.users[id]; !ok {
		return models.ErrUserNotFound
	}

	delete(m.users, id)
	return nil
}

// create a deep copy of user to prevent external modification
func copyUser(user *models.User) *models.User {
	return &models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
