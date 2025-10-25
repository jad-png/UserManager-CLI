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
	username := args[0]
	// TODO: implement login handler logic
}
