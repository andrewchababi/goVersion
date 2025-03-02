package interfaces

import "time"

type Form struct {
	Values map[string]string
	Errors map[string]string
}

type User struct {
	Id           int
	UserName     string
	Password     string
	Organisation string
	CreatedAt    time.Time
}

func NewFormData() Form {
	return Form{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
