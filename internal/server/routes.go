package server

import (
	"revaultier/configuration"
	"revaultier/internal/auth"
	"revaultier/internal/root"
	"revaultier/internal/user"
	"revaultier/internal/vault"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg    *configuration.Config
	Router *echo.Echo
	// params
}

func NewServer(cfg *configuration.Config, rootHandler *root.RootHandler, userHandler *user.UserHandler, authHandler *auth.AuthHandler, vaultHandler *vault.VaultHandler) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtSigningKey := []byte(cfg.Auth.SecretKey)

	e.GET("/", rootHandler.RevaultierStatus, echojwt.JWT(jwtSigningKey))
	e.POST("/login", authHandler.LoginHandler)
	e.POST("/signup", authHandler.SignupHandler)
	e.GET("/vault", vaultHandler.GetVaultsHandler)

	// protected
	r := e.Group("/api")

	return &Server{
		cfg:    cfg,
		Router: e,
	}
}
