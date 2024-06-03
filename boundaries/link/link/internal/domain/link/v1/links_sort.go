package v1

func (l *Links) Len() int {
	return len(l.GetLinks())
}

func (l *Links) Less(i, j int) bool {
	return l.GetLinks()[i].GetCreatedAt().GetTime().Before(l.GetLinks()[j].GetCreatedAt().GetTime())
}

func (l *Links) Swap(i, j int) {
	l.link[i], l.link[j] = l.GetLinks()[j], l.GetLinks()[i]
}
