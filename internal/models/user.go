package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email string, age int) *User {
	now := time.Now()
	return &User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (u *User) Validate() error {
	if u.Name == "" {
		// TODO: return ErrInvalidUser func if no data provided in user creation
	}
	if u.Email == "" {
		// TODO: ErrInvalidUser
	}
	if u.Age < 18 || u.Age > 60 {
		// return error validator for age
	}
	return nil
}

func (u *User) Update(name, email string, age int) {
	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if age >= 0 {
		u.Age = age
	}
	u.UpdatedAt = time.Now()
}
