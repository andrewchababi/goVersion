package api

import (
	"github.com/andrewchababi/restoport/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterUserLoginRoutes(e *echo.Echo) {
	e.POST("/login", handlers.LoginUser)
}
