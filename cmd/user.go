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

func setupUserCommands(commands *commands.UserCommands) {

	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}

	userCmd.AddCommand(
		&cobra.Command{
			Use:   "create [name] [email] [age]",
			Short: "Create a new user",
			Args:  cobra.ExactArgs(3),
			Run:   commands.CreateUser,
		},
		&cobra.Command{
			Use:   "get [id]",
			Short: "Get user by ID",
			Args:  cobra.ExactArgs(1),
			Run:   commands.GetUser,
		},
		&cobra.Command{
			Use:   "list",
			Short: "List all users",
			Run:   commands.GetAllUsers,
		},
		&cobra.Command{
			Use:   "update [id]",
			Short: "Update user information",
			Args:  cobra.ExactArgs(1),
			Run:   commands.UpdateUser,
		},
		&cobra.Command{
			Use:   "delete [id]",
			Short: "Delete a user",
			Args:  cobra.ExactArgs(1),
			Run:   commands.DeleteUser,
		},
	)

	rootCmd.AddCommand(userCmd)
}
