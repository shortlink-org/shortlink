package v1

func (l *Sitemap) Len() int {
	return len(l.GetUrl())
}

func (l *Sitemap) Less(i, j int) bool {
	return l.GetUrl()[i].GetPriority() < l.GetUrl()[j].GetPriority()
}

func (l *Sitemap) Swap(i, j int) {
	l.URL[i], l.URL[j] = l.URL[j], l.URL[i]
}
