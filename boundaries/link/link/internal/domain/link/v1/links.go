package v1

// Link list
type Links struct {
	// Links
	link []*Link
}

// GetLink returns the value of the link field.
func (m *Links) GetLink() []*Link {
	return m.link
}

// Count returns the number of links
func (m *Links) Count() int {
	return len(m.link)
}

func NewLinks() *Links {
	return &Links{}
}

// Push adds a new Link to the link slice
func (l *Links) Push(link ...*Link) {
	l.link = append(l.link, link...)
}
