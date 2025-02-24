package services

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/andrewchababi/restoport/common/interfaces"
	"github.com/andrewchababi/restoport/internal/db"
)

func RunFlightScript() {
	pythonPath := filepath.Join("venv", "Scripts", "python.exe")
	scriptPath := filepath.Join("scripts", "fetch_flights.py")

	cmd := exec.Command(pythonPath, scriptPath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running Python script: %v", err)
	}

	fmt.Println(string(output))
}

func GetCarlosFlights() (interfaces.FlightList, error) {
	flightList, err := db.GetFlightsByGate(62, 68)
	if err != nil {
		return interfaces.FlightList{}, fmt.Errorf("failed to get flights services level: %w", err)
	}
	fmt.Println(flightList.Flights)
	return flightList, nil
}
