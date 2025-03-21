package db

type Env struct {
	DSN       string
	DBName    string
	TableName string
}

func NewEnv() Env {
	return Env{
		DSN:       "root:VavaChab!2!6@tcp(localhost:3306)/flights_data?parseTime=true",
		DBName:    "flights_data",
		TableName: "todays_flights",
	}
}
