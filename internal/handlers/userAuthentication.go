package handlers

import (
	"fmt"

	"github.com/andrewchababi/restoport/common/interfaces"
	"github.com/andrewchababi/restoport/internal/services"
	"github.com/labstack/echo/v4"
)

func InitialPage(c echo.Context) error {
	message := "Login"
	form := interfaces.NewFormData()

	page := interfaces.NewPage(message, form)
	return c.Render(200, "userauth", page)
}

func LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Printf("Recieved Data -> Name: %v Email: %v \n", username, password)

	organisation, err := services.UserAuth(username, password)

	if err != nil {
		return c.HTML(200, fmt.Sprintf(`<div style="color:red;">Error: %s</div>`, err.Error()))
	}

	services.Redirect(c, organisation)
	return nil
}
