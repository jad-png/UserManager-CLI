package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "user-manager",
	Short: "userManager is a CLI for managing users in memory",
	Long:  `UserManager CLI - A learning project for Go A complete CRUD application with authentication and concurrency features built for learning Go fundamentals`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the UserManager CLI")
		fmt.Println("Use 'user-manager --help' to see available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Subcommand definitions
var (
	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Authentication operations",
		Long:  "Manage authentication and user sessions",
	}

	userCmd = &cobra.Command{
		Use:   "user",
		Short: "User management",
		Long:  "Create, read, update, and delete users",
	}

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configuration management",
		Long:  "View and modify application configuration",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("UserManager CLI v0.1.0")
		},
	}
)

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP(" ", "c", "", "config file (default is $HOME/.user-manager.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// Setup subcommands first
	setupSubcommands()

	// Add subcommands to root
	rootCmd.AddCommand(authCmd, userCmd, configCmd, versionCmd)
}

func setupSubcommands() {
	// User subcommands - using simple commands for now
	userCmd.AddCommand(
		&cobra.Command{
			Use:   "create [name] [email] [age]",
			Short: "Create a new user",
			Args:  cobra.ExactArgs(3),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Creating user: %s, %s, %s\n", args[0], args[1], args[2])
			},
		},
		&cobra.Command{
			Use:   "get [id]",
			Short: "Get user by ID",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Getting user: %s\n", args[0])
			},
		},
		&cobra.Command{
			Use:   "list",
			Short: "List all users",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Listing all users")
			},
		},
		&cobra.Command{
			Use:   "update [id]",
			Short: "Update user information",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Updating user: %s\n", args[0])
			},
		},
		&cobra.Command{
			Use:   "delete [id]",
			Short: "Delete a user",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Deleting user: %s\n", args[0])
			},
		},
	)

	// Auth subcommands
	authCmd.AddCommand(
		&cobra.Command{
			Use:   "login [username]",
			Short: "Login to the system",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Logging in user: %s\n", args[0])
			},
		},
		&cobra.Command{
			Use:   "logout",
			Short: "Logout from the system",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Logging out")
			},
		},
		&cobra.Command{
			Use:   "status",
			Short: "Show authentication status",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Authentication status: Not logged in")
			},
		},
	)
}
