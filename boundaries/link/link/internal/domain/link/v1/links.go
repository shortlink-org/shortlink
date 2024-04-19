package v1

func NewLinks() *Links {
	return &Links{}
}

// Push adds a new Link to the link slice
func (l *Links) Push(link ...*Link) {
	l.link = append(l.link, link...)
}
