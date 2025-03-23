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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("home"))
	})
	e.GET("/home", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("home"))
	})
	e.GET("/music", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("music"))
	})
	e.GET("/list", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("list"))
	})
	e.GET("/settings", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("settings"))
	})
	e.GET("/account", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page("account"))
	})
	e.GET("/views/home", func(c echo.Context) error {
		return render(c, http.StatusOK, views.HomePage())
	})
	e.GET("/views/music", func(c echo.Context) error {
		return render(c, http.StatusOK, views.MusicPage())
	})
	e.GET("/views/list", func(c echo.Context) error {
		return render(c, http.StatusOK, views.ListPage())
	})
	e.GET("/views/settings", func(c echo.Context) error {
		return render(c, http.StatusOK, views.SettingsPage(data.Settings()))
	})
	e.GET("/views/account", func(c echo.Context) error {
		return render(c, http.StatusOK, views.AccountPage())
	})
	staticHandler := http.FileServer(http.FS(static))
	e.GET("/static/*", echo.WrapHandler(staticHandler))
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

func render(c echo.Context, status int, t templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	err := t.Render(context.Background(), c.Response().Writer)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to render component")
	}
	return nil
}
