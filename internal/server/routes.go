package server

import (
	"revaultier/internal/auth"
	"revaultier/internal/root"
	"revaultier/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(rootHandler *root.RootHandler, userHandler *user.UserHandler, authHandler *auth.AuthHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler.RevaultierStatus)
	e.POST("/login", authHandler.LoginHandler)
	e.POST("/signup", authHandler.SignupHandler)

	return e
}
