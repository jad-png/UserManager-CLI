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

func (uc *UserCommands) GetUser(cmd *cobra.Command, args []string) {
	id := args[0]

	user, err := uc.Storage.GetById(id)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User found:\n")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("Created: %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", user.UpdatedAt.Format("2006-01-02 15:04:05"))
}

func (uc *UserCommands) GetAllUsers(cmd *cobra.Command, args []string) {
	users, err := uc.Storage.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(users) == 0 {
		fmt.Printf("No users found\n")
		return
	}

	fmt.Printf("Found %d users:\n", len(users))
	fmt.Println("=========================================")
	for i, user := range users {
		fmt.Printf("%d. ID: %s\n", i+1, user.ID)
		fmt.Printf("   Name: %s\n", user.Name)
		fmt.Printf("   Email: %s\n", user.Email)
		fmt.Printf("   Age: %d\n", user.Age)
		fmt.Printf("   Created: %s\n", user.CreatedAt.Format("2006-01-02"))
		fmt.Println("-----------------------------------------")
	}
}
