package interfaces

type Page struct {
	Title   string
	Content PageContent
}

type PageContent interface {
}

func NewPage(title string, content PageContent) Page {
	return Page{
		Title:   title,
		Content: content,
	}
}
