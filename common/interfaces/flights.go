package interfaces

type Flight struct {
	AirLineName     string
	Gate            int
	Time            string
	UpdatedTime     string
	AirportName     string
	Status          string
	UniqueDisplayNo string
}

type FlightList struct {
	Flights []Flight
}
