package handlers

import (
	"fmt"

	"github.com/andrewchababi/restoport/common/interfaces"
	"github.com/andrewchababi/restoport/internal/services"
	"github.com/labstack/echo/v4"
)

func InitialPage(c echo.Context) error {
	message := "Welcome to Carlos & Pepes flights"
	flights, err := services.GetCarlosFlights()
	if err != nil {
		return fmt.Errorf("could not load flights handler level %w", err)
	}

	page := interfaces.NewPage(message, flights)
	return c.Render(200, "index", page)
}
