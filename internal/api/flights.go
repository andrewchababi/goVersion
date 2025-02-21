package api

import (
	"github.com/andrewchababi/restoport/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterFlightsRoutes(e *echo.Echo) {
	e.GET("/", handlers.InitialPage)
}
