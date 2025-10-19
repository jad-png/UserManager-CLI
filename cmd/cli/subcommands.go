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
		Run:   CreateUser,
	}
	userCreateCmd.Flags().StringP("role", "r", "user", "User role")

	userGetCmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get user by ID",
		Args:  cobra.ExactArgs(1),
		Run:   GetUser,
	}

	userListCmd := &cobra.Command{
		Use:   "list",
		Short: "List all users",
		Run:   ListUsers,
	}
	userListCmd.Flags().IntP("limit", "l", 50, "Maximum users to display")

	userUpdateCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update user information",
		Args:  cobra.ExactArgs(1),
		Run:   UpdateUser,
	}
	userUpdateCmd.Flags().StringP("name", "n", "", "New name")
	userUpdateCmd.Flags().StringP("email", "e", "", "New email")
	userUpdateCmd.Flags().IntP("age", "a", 0, "New age")

	userDeleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a user",
		Args:  cobra.ExactArgs(1),
		Run:   DeleteUser,
	}

	// Add user subcommands to user command
	getUserCommand().AddCommand(userCreateCmd, userGetCmd, userListCmd, userUpdateCmd, userDeleteCmd)

	setupAuthCommands()
}

// helper method to get command (to be linked from root.go)
func getUserCommand() *cobra.Command {
	return &cobra.Command{Use: "user"}
}

func setupAuthCommands() {
	authLoginCmd := &cobra.Command{
		Use:   "Login [username]",
		Short: "Login to the system",
		Args:  cobra.ExactArgs(1),
		Run:   Login,
	}

	authLogoutCmd := &cobra.Command{
		Use:   "Logout",
		Short: "Logout the system",
		Run:   Logout,
	}

	getAuthCommand().AddCommand(authLoginCmd, authLogoutCmd)
}

func getAuthCommand() *cobra.Command {
	return &cobra.Command{Use: "auth"} // placeholder
}

// TODO: function to be implemented later
func CreateUser(cmd *cobra.Command, args []string) {
	fmt.Printf("Creating user: %s, %s, %s\n", args[0], args[1], args[2])
}

func GetUser(cmd *cobra.Command, args []string) {
	fmt.Printf("Getting user: %s\n", args[0])
}

func ListUsers(cmd *cobra.Command, args []string) {
	fmt.Println("Listing all users")
}

func UpdateUser(cmd *cobra.Command, args []string) {
	fmt.Printf("Updating user: %s\n", args[0])
}

func DeleteUser(cmd *cobra.Command, args []string) {
	fmt.Printf("Deleting user: %s\n", args[0])
}

func Login(cmd *cobra.Command, args []string) {
	fmt.Printf("Logging in user: %s\n", args[0])
}

func Logout(cmd *cobra.Command, args []string) {
	fmt.Println("Logging out")
}
