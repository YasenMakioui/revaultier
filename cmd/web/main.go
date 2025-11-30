package main

import (
	"revaultier/configuration"
	"revaultier/database"
	"revaultier/internal/auth"
	"revaultier/internal/root"
	"revaultier/internal/server"
	"revaultier/internal/user"
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

	e := server.NewServer(cfg, rootHandler, userHandler, authHandler)

	if err := e.Router.Start(":8080"); err != nil {
		e.Router.Logger.Fatal("could not start server: ", err)
	}
}
