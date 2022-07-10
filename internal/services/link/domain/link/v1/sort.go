package v1

func (l *Links) Len() int {
	return len(l.Link)
}

func (l *Links) Less(i, j int) bool {
	return l.Link[i].CreatedAt.AsTime().Before(l.Link[j].CreatedAt.AsTime())
}

func (l *Links) Swap(i, j int) {
	l.Link[i], l.Link[j] = l.Link[j], l.Link[i]
}
