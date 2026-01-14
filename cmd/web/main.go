package main

import (
	"revaultier/configuration"
	"revaultier/database"
	"revaultier/internal/auth"
	"revaultier/internal/card"
	"revaultier/internal/root"
	"revaultier/internal/server"
	"revaultier/internal/user"
	"revaultier/internal/vault"
)

func main() {

	cfg, err := configuration.LoadConfig()

	if err != nil {
		panic(err)
	}

	database := database.NewDatabase(cfg)

	rootHandler := root.NewRootHandler()

	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(cfg, userRepository)
	userHandler := user.NewUserHandler(userService)

	authRepository := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(cfg, authRepository)
	authHandler := auth.NewAuthHandler(authService)

	vaultRepository := vault.NewVaultRepository(database)
	vaultService := vault.NewVaultService(cfg, vaultRepository)
	vaultHandler := vault.NewVaultHandler(vaultService)

	cardRepository := card.NewCardRepository(database)
	cardService := card.NewCardService(cfg, cardRepository)
	cardHandler := card.NewCardHandler(cardService)

	e := server.NewServer(cfg, rootHandler, userHandler, authHandler, vaultHandler, cardHandler)

	if err := e.Router.Start(":8080"); err != nil {
		e.Router.Logger.Fatal("could not start server: ", err)
	}
}
