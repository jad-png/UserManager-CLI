package cmd

import (
	"awesomeProject/internal/commands"

	"github.com/spf13/cobra"
)

func init() {
	authCommands := commands.NewAuthCommands()

	setupAuthCommands(authCommands)
}

func setupAuthCommands(authCommands *commands.AuthCommands) {
	authCmd.AddCommand(
		&cobra.Command{
			Use:   "login [username]",
			Short: "Login to the system",
			Args:  cobra.ExactArgs(1),
			Run:   authCommands.Login,
		},
		&cobra.Command{
			Use:   "logout",
			Short: "Logout from the system",
			Run:   authCommands.Logout,
		},
	)

	rootCmd.AddCommand(authCmd)
}
