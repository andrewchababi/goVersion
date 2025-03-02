package handlers

import (
	"database/sql"
	"errors"
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

	_, err := services.UserAuth(username, password)

	if err != nil {
		// If no user was found or another error occurred, return an HTML snippet with the error message.
		if errors.Is(err, sql.ErrNoRows) {
			return c.HTML(200, `<div style="color:red;">User does not exist</div>`)
		}
		return c.HTML(200, fmt.Sprintf(`<div style="color:red;">Error: %s</div>`, err.Error()))
	}

	// If authentication is successful, you might want to render a dashboard page or similar.
	// For this example, we just return the organisation.

	c.Response().Header().Set("HX-Redirect", "/carlos")
	return nil
}
