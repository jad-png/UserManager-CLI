package commands

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type UserCommands struct {
	Storage storage.Storage
}

func NewUserCommands(Storage storage.Storage) *UserCommands {
	return &UserCommands{
		Storage: Storage,
	}
}

func (uc *UserCommands) createUser(cmd *cobra.Command, args []string) {
	name := args[0]
	email := args[1]
	age, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error: age must be a number - %v\n", err)
		return
	}

	user := models.NewUser(name, email, age)

	if err := uc.Storage.Create(user); err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}

	fmt.Printf("User created successfully!\n")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Age: %d\n", user.Age)
}
