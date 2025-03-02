package db

import (
	"database/sql"
	"fmt"

	"github.com/andrewchababi/restoport/common/interfaces"
)

func UserExist(username string) (*interfaces.User, error) {
	query := "SELECT id, username, password, organisation, created_at FROM users WHERE username = ? LIMIT 1"
	row := db.QueryRow(query, username)

	var user interfaces.User

	if err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Organisation, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error at db service level. No user was found: %w", err) // No user found
		}
		return nil, err // Some other error occurred
	}
	return &user, nil
}
