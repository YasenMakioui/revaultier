package main

import (
	"fmt"
	"log"
	"revaultier/database"
	"revaultier/internal/auth"
	"revaultier/internal/root"
	"revaultier/internal/server"
	"revaultier/internal/user"

	"github.com/spf13/viper"
)

func main() {

	// Get config file. Move this in internal to be retrieved by main
	viper.SetConfigFile("revaultier.toml")
	viper.AddConfigPath("$HOME/.config/revaultier")
	viper.AddConfigPath("/etc/revaultier")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}

	secretkey := []byte(viper.GetString("auth.secretkey"))
	databaseName := viper.GetString("database.database")

	if string(secretkey) == "" {
		log.Fatal(fmt.Errorf("set the secretkey in the configuration file"))
	}

	database := database.NewDatabase(databaseName)

	rootHandler := root.NewRootHandler()

	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	authRepository := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(authRepository, secretkey) // temp, will be taken from config
	authHandler := auth.NewAuthHandler(authService)

	e := server.NewServer([]byte("secretkey"), rootHandler, userHandler, authHandler)

	if err := e.Router.Start(":8080"); err != nil {
		e.Router.Logger.Fatal("could not start server: ", err)
	}
}
