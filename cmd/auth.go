package cmd

import (
	"awesomeProject/internal/commands"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	authCommands := commands.NewAuthCommands()
}

func setupAuthCommands(authCommands *commands.AuthCommands) {
	authCmd.AddCommand()
}
