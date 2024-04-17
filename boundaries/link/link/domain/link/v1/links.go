package v1

func NewLinks() *Links {
	return &Links{}
}

// Push - push link to list
func (l *Links) Push(link *Link) {
	l.link = append(l.GetLink(), link)
}
