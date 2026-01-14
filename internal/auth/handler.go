package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(as *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: as}
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {

	ctx := c.Request().Context()

	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid payload"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "username and password are required"})
	}

	if err := h.AuthService.AuthenticateService(ctx, req.Username, req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "failed authentication"})
	}

	uuid, err := h.AuthService.GetUserUUIDService(ctx, req.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not authenticate"})
	}

	token, err := h.AuthService.GenerateTokenSerivce(req.Username, uuid) // should pass the uuid and add it in the "sub" claim

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not authenticate"})
	}

	return c.JSON(http.StatusOK, AuthResponse{Token: token})
}

func (h *AuthHandler) SignupHandler(c echo.Context) error {

	ctx := c.Request().Context()

	req := new(SignupRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid payload"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "username and password are required"})
	}

	if err := h.AuthService.SignupService(ctx, req.Username, req.Password); err != nil {
		switch err {
		case ErrUsernameTaken:
			return c.JSON(http.StatusConflict, ErrorResponse{Error: "username already taken"})
		default:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not create user"})
		}
	}

	return c.JSON(http.StatusOK, SignupResponse{Username: req.Username})
}
