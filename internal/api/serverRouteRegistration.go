package api

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	RegisterDashBoardRoutes(e)
	RegisterHomePageRoutes(e)
	RegisterUserLoginRoutes(e)
}
