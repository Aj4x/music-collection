package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Aj4x/music-collection/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Page())
	})
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
