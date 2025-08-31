package server

import (
	"html/template"
	"io"

	"github.com/YasenMakioui/revaultier/internal/root"
	"github.com/YasenMakioui/revaultier/internal/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewRouter(rootHandler *root.RootHandler, userHandler *user.UserHandler) *echo.Echo {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}

	e.Renderer = t

	// e.GET("/login", userHandler.ShowLoginPage)
	// e.POST("/login", userHandler.Authenticate)
	e.GET("/", rootHandler.ShowRootPage)
	e.POST("/", rootHandler.ShowRootPage)
	//e.GET("/revaultier/cards", cardHandler.ShowCardsPage)
	e.GET("/signup", userHandler.ShowSignupPage)
	e.POST("/signup", userHandler.Signup)
	e.GET("/login", userHandler.ShowLoginPage)
	e.POST("/login", userHandler.Login)
	e.GET("/logout", userHandler.Logout)

	return e
}
