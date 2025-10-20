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

func (s *Service) CreateUser(name, email string, age int) (*models.User, error) {
	user := models.NewUser(name, email, age)

	if err := s.store.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUser(id string) (*models.User, error) {
	return s.store.GetById(id)
}

func (s *Service) GetAll() ([]*models.User, error) {
	return s.store.GetAll()
}

func (s *Service) Update(id string, name, email string, age int) (*models.User, error) {

	user, err := s.store.GetById(id)
	if err != nil {
		return nil, err
	}

	updatedUser := &models.User{
		ID:        user.ID,
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if ok := s.store.Update(id, updatedUser); ok != nil {
		return nil, ok
	}

	return s.store.GetById(id)
}

func (s *Service) Delete(id string) error {
	return s.store.Delete(id)
}

func (s *Service) Exists(id string) bool {
	return s.store.Exists(id)
}

func (s *Service) Count() int {
	return s.store.Count()
}
