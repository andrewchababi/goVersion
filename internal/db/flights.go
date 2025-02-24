package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/andrewchababi/restoport/common/interfaces"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	createDBInstance()
}

func createDBInstance() error {
	env := NewEnv()
	var err error

	db, err = sql.Open("mysql", env.DSN)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}

func GetFlightsByGate(g1 int, g2 int) (interfaces.FlightList, error) {
	query := "SELECT * FROM todays_flights WHERE gate BETWEEN ? AND ?"

	rows, err := db.Query(query, g1, g2)
	if err != nil {
		return interfaces.FlightList{}, fmt.Errorf("failed to querry flights: %w", err)
	}
	defer rows.Close()

	var flights interfaces.FlightList

	for rows.Next() {
		var flight interfaces.Flight
		err := rows.Scan(&flight.AirLineName, &flight.Gate, &flight.Time, &flight.UpdatedTime, &flight.AirportName, &flight.Status, &flight.UniqueDisplayNo)
		if err != nil {
			return interfaces.FlightList{}, fmt.Errorf("failed to scan row: %w", err)
		}

		flights.Flights = append(flights.Flights, flight)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return interfaces.FlightList{}, fmt.Errorf("row iteration error: %w", err)
	}

	return flights, nil
}
