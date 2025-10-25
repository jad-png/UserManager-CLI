package commands

import "github.com/spf13/cobra"

type AuthCommands struct {
	// auth dependencies
}

// new instance
func NewAuthCommands() *AuthCommands {
	return &AuthCommands{}
}

func (ac *AuthCommands) Login(cmd *cobra.Command, args []string) {
	//name := args[0]
	// TODO: implement login handler logic
}

func (ac *AuthCommands) Logout(cmd *cobra.Command, args []string) {
	// TODO: implement logout
}
