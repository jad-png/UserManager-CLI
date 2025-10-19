package models

import "fmt"

type ErrInvalidUser struct {
	Field  string
	Reason string
}

func (e ErrInvalidUser) Error() string {
	return fmt.Sprintf("Invalide user %s: %s", e.Field, e.Reason)
}
