package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Age          int       `json:"age"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
		return ErrInvalidUser{Field: "name", Reason: "cannot be empty"}
	}
	if u.Email == "" {
		return ErrInvalidUser{Field: "email", Reason: "cannot be empty"}
	}
	if u.Age < 0 || u.Age > 150 {
		return ErrInvalidUser{Field: "age", Reason: "must be between 0 and 150"}
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

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}
