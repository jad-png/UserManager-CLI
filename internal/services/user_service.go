package services

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
)

type Service struct {
	userStorage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		userStorage: storage,
	}
}

func (s *Service) CreateUser(name, email string, age int) (*models.User, error) {
	user := models.NewUser(name, email, age)

	if err := s.userStorage.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUser(id string) (*models.User, error) {
	return s.userStorage.GetById(id)
}

func (s *Service) GetAll() ([]*models.User, error) {
	return s.userStorage.GetAll()
}

func (s *Service) Update(id string, name, email string, age int) (*models.User, error) {

	user, err := s.userStorage.GetById(id)
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

	if ok := s.userStorage.Update(id, updatedUser); ok != nil {
		return nil, ok
	}

	return s.userStorage.GetById(id)
}

func (s *Service) Delete(id string) error {
	return s.userStorage.Delete(id)
}

func (s *Service) Exists(id string) bool {
	return s.userStorage.Exists(id)
}

func (s *Service) Count() int {
	return s.userStorage.Count()
}
