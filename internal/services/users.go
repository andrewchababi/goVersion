package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/andrewchababi/restoport/internal/db"
)

func UserAuth(username, password string) (string, error) {
	user, err := db.UserExist(username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("no user was found with such username: %w", err)
		}
		return "", fmt.Errorf("issue retrieving user data: %w", err)
	}

	if user.Password != password {
		return "", fmt.Errorf("wrong password")
	}

	return user.Organisation, nil
}

func Redirect(organisation string) {

}
