package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

import "awesomeProject/cmd/cli"

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

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.user-manager.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// TODO: add subcommands
	// exemple: rootCmd.AddCommand(ExempleCmd)
}

// TODO: define Subcommands in rootCmd style

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication operations",
	Long:  "Manage authentication and user sessions",
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
	Long:  "Create, read, update, and delete users",
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
	Long:  "View and modify application configuration",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UserManager CLI v0.1.0")
	},
}

func setupSubcommands() {
	// User subcommands
	userCmd.AddCommand(
		&cobra.Command{Use: "create",
			Short: "Create user",
			Run:   cli.CreateUser,
		},

		&cobra.Command{Use: "get",
			Short: "Get user",
			Run:   cli.GetUser,
		},
		&cobra.Command{Use: "list",
			Short: "List users",
			Run:   cli.ListUsers,
		},
		&cobra.Command{Use: "update",
			Short: "Update user",
			Run:   cli.UpdateUser,
		},
		&cobra.Command{Use: "delete",
			Short: "Delete user",
			Run:   cli.DeleteUser,
		},
	)

	// Auth subcommands
	authCmd.AddCommand(
		&cobra.Command{Use: "login",
			Short: "Login",
			Run:   cli.Login,
		},
		&cobra.Command{Use: "logout",
			Short: "Logout",
			Run:   cli.Logout,
		},
	)
}
