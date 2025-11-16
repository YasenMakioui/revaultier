package root

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

// func (r *RootHandler) ShowRootPage(c echo.Context) error {
// 	return c.Render(http.StatusOK, "index.html", map[any]any{})
// }

func (h *RootHandler) RevaultierStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"version": "v1.0.0",
		"status":  "ok",
		"about":   "you know, to be secure",
	})
}
