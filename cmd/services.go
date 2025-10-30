package cmd

import (
	"awesomeProject/internal/commands"
	"awesomeProject/internal/services"
	"awesomeProject/internal/storage"
)

var (
	userStorage = storage.NewMemory()
	authService = services.NewAuthService(userStorage)

	userCommands = commands.NewUserCommands(userStorage, authService)
	authCommands = commands.NewAuthCommands(authService)
)
