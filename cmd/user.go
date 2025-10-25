package cmd

import (
	"awesomeProject/internal/commands"
	"awesomeProject/internal/storage"

	"github.com/spf13/cobra"
)

func init() {
	userStorage := storage.NewMemory()
	userCommands := commands.NewUserCommands(userStorage)

	setupUserCommands(userCommands)
}

func setupUserCommands(userCommands *commands.UserCommands) {

	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}

	userCmd.AddCommand(
		&cobra.Command{
			Use:   "create [name] [email] [age]",
			Short: "Create a new user",
			Args:  cobra.ExactArgs(3),
			Run:   userCommands.CreateUser,
		},
		&cobra.Command{
			Use:   "get [id]",
			Short: "Get user by ID",
			Args:  cobra.ExactArgs(1),
			Run:   userCommands.GetUser,
		},
		&cobra.Command{
			Use:   "list",
			Short: "List all users",
			Run:   userCommands.GetAllUsers,
		},
		&cobra.Command{
			Use:   "update [id]",
			Short: "Update user information",
			Args:  cobra.ExactArgs(1),
			Run:   userCommands.UpdateUser,
		},
		&cobra.Command{
			Use:   "delete [id]",
			Short: "Delete a user",
			Args:  cobra.ExactArgs(1),
			Run:   userCommands.DeleteUser,
		},
	)

	rootCmd.AddCommand(userCmd)
}
