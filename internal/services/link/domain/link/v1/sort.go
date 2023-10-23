package v1

func (l *Links) Len() int {
	return len(l.GetLink())
}

func (l *Links) Less(i, j int) bool {
	return l.GetLink()[i].GetCreatedAt().AsTime().Before(l.GetLink()[j].GetCreatedAt().AsTime())
}

func (l *Links) Swap(i, j int) {
	l.Link[i], l.Link[j] = l.GetLink()[j], l.GetLink()[i]
}
