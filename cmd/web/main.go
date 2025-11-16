package main

import (
	"revaultier/database"
	"revaultier/internal/auth"
	"revaultier/internal/root"
	"revaultier/internal/server"
	"revaultier/internal/user"
)

func main() {
	database := database.NewDatabase("./test.db")

	rootHandler := root.NewRootHandler()

	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	authRepository := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(authRepository, []byte("secretkey")) // temp, will be taken from config
	authHandler := auth.NewAuthHandler(authService)

	e := server.NewRouter(rootHandler, userHandler, authHandler)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("could not start server: ", err)
	}
}
