package main

import (
	"github.com/YasenMakioui/revaultier/database"
	"github.com/YasenMakioui/revaultier/internal/root"
	"github.com/YasenMakioui/revaultier/internal/server"
	"github.com/YasenMakioui/revaultier/internal/user"
)

func main() {
	// Setup database
	database := database.NewDatabase("./revaultier.db")

	rootHandler := root.NewHandler()

	userRepository := user.NewRepository(database)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	e := server.NewRouter(rootHandler, userHandler)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("could not start server: ", err)
	}

}
