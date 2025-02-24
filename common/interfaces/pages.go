package interfaces

type Page struct {
	Title      string
	FlightList FlightList
}

func NewPage(title string, flights FlightList) Page {
	return Page{
		Title:      title,
		FlightList: flights,
	}
}
