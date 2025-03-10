package api

import (
	"github.com/andrewchababi/restoport/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterHomePageRoutes(e *echo.Echo) {
	e.GET("/home", handlers.HomePage)
}
