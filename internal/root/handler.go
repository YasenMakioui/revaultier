package root

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type RootHandler struct{}

func NewHandler() *RootHandler {
	return &RootHandler{}
}

func (r *RootHandler) ShowRootPage(c echo.Context) error {
	loggedIn := false

	sess, err := session.Get("session", c)

	if err != nil {
		return c.Render(http.StatusInternalServerError, "internalServerError.html", map[string]any{})
	}

	if sess.Values["username"] != nil {
		loggedIn = true
	}

	log.Println(loggedIn)

	return c.Render(http.StatusOK, "index.html", map[string]any{
		"username": sess.Values["username"],
		"loggedIn": loggedIn,
	})
}
