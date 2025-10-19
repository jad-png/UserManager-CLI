package storage

import "awesomeProject/internal/models"

type Storage interface {
	Create(user *models.User) error

	GetById(id string) (*models.User, error)

	GetByEmail(email string) (*models.User, error)

	GetAll() ([]*models.User, error)

	Update(id string, updatedUser *models.User) error

	Delete(id string) error

	Exists(id string) bool

	Count() int
}
