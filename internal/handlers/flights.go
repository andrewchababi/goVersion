package handlers

import (
	"github.com/andrewchababi/restoport/internal/interfaces"
	"github.com/labstack/echo/v4"
)

func InitialPage(c echo.Context) error {
	page := interfaces.NewPage("Welcome")
	return c.Render(200, "index", page)
}
