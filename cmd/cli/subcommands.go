package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	userCreateCmd := &cobra.Command{
		Use:   "create [name] [email] [age]",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(3),
		//	Run:   call createUser func,
	}
	userCreateCmd.Flags().StringP("role", "r", "user", "User role")

	userGetCmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get user by ID",
		Args:  cobra.ExactArgs(1),
		//Run:   getUser,
	}

	userListCmd := &cobra.Command{
		Use:   "list",
		Short: "List all users",
		//Run:   listUsers,
	}
	userListCmd.Flags().IntP("limit", "l", 50, "Maximum users to display")

	userUpdateCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update user information",
		Args:  cobra.ExactArgs(1),
		//Run:   updateUser,
	}
	userUpdateCmd.Flags().StringP("name", "n", "", "New name")
	userUpdateCmd.Flags().StringP("email", "e", "", "New email")
	userUpdateCmd.Flags().IntP("age", "a", 0, "New age")

	userDeleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a user",
		Args:  cobra.ExactArgs(1),
		//Run:   deleteUser,
	}

	// Add user subcommands to user command
}

// helper method to get command (to be linked from root.go)
func getUserCommand() *cobra.Command {
	return &cobra.Command{Use: "user"}
}

func setupAuthCommands() {
	authLoginCmd := &cobra.Command{
		Use: "Login [username]";
		Short: "Login to the system",
		Args:  cobra.ExactArgs(1),
		//Run: Login
	}

	authLogoutCmd := &cobra.Command{
		Use: "Logout",
		Short: "Logout the system",
		//Run: Logout
	}
}

func getAuthCommand() *cobra.Command {
	return &cobra.Command{Use: "auth"} // placeholder
}