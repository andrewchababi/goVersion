package main

import (
	"github.com/andrewchababi/restoport/internal/api"
	"github.com/andrewchababi/restoport/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/web", "web")

	e.Renderer = services.NewTemplates()

	api.RegisterFlightsRoutes(e)

	services.RunFlightScript()

	e.Logger.Fatal(e.Start(":5000"))
}
