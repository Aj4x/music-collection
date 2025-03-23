package main

import (
	"context"
	"embed"
	"errors"
	"github.com/Aj4x/music-collection/data"
	"log/slog"
	"net/http"

	"github.com/Aj4x/music-collection/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static
var static embed.FS

func main() {
	// new echo server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// page routing
	e.GET("/", renderView(views.Page("home")))
	e.GET("/home", renderView(views.Page("home")))
	e.GET("/music", renderView(views.Page("music")))
	e.GET("/list", renderView(views.Page("list")))
	e.GET("/settings", renderView(views.Page("settings")))
	e.GET("/account", renderView(views.Page("account")))

	// partial views for htmx
	e.GET("/views/home", renderView(views.HomePage()))
	e.GET("/views/music", renderView(views.MusicPage()))
	e.GET("/views/list", renderView(views.ListPage()))
	e.GET("/views/settings", renderView(views.SettingsPage(data.Settings())))
	e.GET("/views/account", renderView(views.AccountPage()))

	//	static file server
	staticHandler := http.FileServer(http.FS(static))
	e.GET("/static/*", echo.WrapHandler(staticHandler))

	//start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

func renderView(t templ.Component) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Writer.WriteHeader(http.StatusOK)
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		err := t.Render(context.Background(), c.Response().Writer)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to render component")
		}
		return nil
	}
}
