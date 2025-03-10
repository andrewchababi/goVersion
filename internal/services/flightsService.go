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

func GetFlightsByOrganisation(org string) (*interfaces.FlightList, error) {
	g1, g2, err := db.GetGateMap(org)
	if err != nil {
		return nil, fmt.Errorf("no gates maped to %s organisation", org)
	}

	flightList, err := db.GetFlightsByGate(g1, g2)
	if err != nil {
		return nil, fmt.Errorf("failed to get flights services level: %w", err)
	}
	return flightList, nil
}
