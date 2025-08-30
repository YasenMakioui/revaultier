package user

import (
	"net/http"

	"github.com/YasenMakioui/revaultier/internal/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *UserService
}

func NewHandler(us *UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) ShowLoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]any{
		"title": "Login",
	})
}

func (h *UserHandler) ShowSignupPage(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", map[string]any{
		"title": "Signup",
	})
}

func (h *UserHandler) Signup(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate credentials

	if err := utils.ValidatePassword(password); err != nil {
		return c.Render(http.StatusOK, "signupError.html", map[string]any{"message": err.Error(), "passwordFormat": true})
	}

	// Validate if user exists

	if h.userService.UserExists(email) {
		return c.Render(http.StatusOK, "signupError.html", map[string]any{"message": "User Exists"})
	}

	if err := h.userService.Signup(email, password); err != nil {
		return c.Render(http.StatusOK, "signupError.html", map[string]any{"message": "Error while signing up"})
	}

	return c.Render(http.StatusOK, "success.html", map[string]any{})
}

func (h *UserHandler) Login(c echo.Context) error {
	// Shouls check if the user and password are correct
	// give a session token with echo
	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := h.userService.Authenticate(username, password); err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	// create session
	if err := h.userService.CreateSession(c, username); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// func (h *Handler) Authenticate(c echo.Context) error {
// 	// Get username and signature
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	if err := h.userService.Authenticate(email, password); err != nil {
// 		return c.String(http.StatusForbidden, "bad credentials")
// 	}
// 	return c.String(http.StatusOK, "okiiii")
// }
