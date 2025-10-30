package commands

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

	"golang.org/x/term"
	"os"
	"syscall"
)

type UserCommands struct {
	Storage storage.Storage
}

func NewUserCommands(Storage storage.Storage) *UserCommands {
	return &UserCommands{
		Storage: Storage,
	}
}

func (uc *UserCommands) CreateUser(cmd *cobra.Command, args []string) {
	name := args[0]
	email := args[1]
	age, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Printf("Error: age must be a number - %v\n", err)
		return
	}

	// --- block to securely read pswrd ---
	fmt.Print("Enter Password:")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Error reading password: %v\n", err)
		return
	}
	password := string(bytePassword)
	fmt.Println()

	// -- password confirmation
	fmt.Print("Confirm Password:")
	bytePasswordConfirm, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Error reading password: %v\n", err)
		return
	}
	passwordConfirm := string(bytePasswordConfirm)
	fmt.Println()

	if password != passwordConfirm {
		fmt.Println("Passwords do not match")
		return
	}
	// --- end of block ---

	user, err := models.NewUser(name, email, age, password)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}

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

func (uc *UserCommands) UpdateUser(cmd *cobra.Command, args []string) {
	id := args[0]

	name, _ := cmd.Flags().GetString("name")
	email, _ := cmd.Flags().GetString("email")
	age, _ := cmd.Flags().GetInt("age")

	existingUser, err := uc.Storage.GetById(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if name == "" {
		name = existingUser.Name
	}
	if email == "" {
		email = existingUser.Email
	}
	if age == 0 {
		age = existingUser.Age
	}

	updatedUser := &models.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	if err := uc.Storage.Update(id, updatedUser); err != nil {
		fmt.Printf("Error updating: %v\n", err)
		return
	}

	fmt.Printf("User updated successfully!\n")
}

func (uc *UserCommands) DeleteUser(cmd *cobra.Command, args []string) {
	id := args[0]

	if err := uc.Storage.Delete(id); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User deleted successfully!\n")
}
