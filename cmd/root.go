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

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.user-manager.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// TODO: add subcommands
	// exemple: rootCmd.AddCommand(ExempleCmd)
}
