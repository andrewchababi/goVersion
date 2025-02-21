package interfaces

type Page struct {
	Title string
}

func NewPage(title string) Page {
	return Page{
		Title: title,
	}
}
