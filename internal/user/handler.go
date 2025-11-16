package user

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(us *UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

// func (h *UserHandler) ShowLoginPage(c echo.Context) error {
// 	return c.Render(http.StatusOK, "login.html", map[string]any{})
// }
