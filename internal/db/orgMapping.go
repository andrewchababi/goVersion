package db

import "fmt"

func GetGateMap(org string) (int, int, error) {
	querry := "SELECT lower_gate, upper_gate FROM organisation_gate_map WHERE organisation = ? LIMIT 1"
	row := db.QueryRow(querry, org)

	var g1, g2 int

	if err := row.Scan(&g1, &g2); err != nil {
		return 0, 0, fmt.Errorf("problem retreiving gates for organisation")
	}

	return g1, g2, nil
}
