package handlers

import (
	"fmt"

	"github.com/andrewchababi/restoport/common/interfaces"
	"github.com/andrewchababi/restoport/internal/services"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	org := c.QueryParam("org")

	message := "Welcome to RestoPort"
	flights, err := services.GetFlightsByOrganisation(org)
	if err != nil {
		return fmt.Errorf("could not load flights handler level %w", err)
	}

	page := interfaces.NewPage(message, flights)
	return c.Render(200, "index", page)
}
