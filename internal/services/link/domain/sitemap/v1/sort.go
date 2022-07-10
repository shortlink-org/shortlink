package v1

func (l *Sitemap) Len() int {
	return len(l.Url)
}

func (l *Sitemap) Less(i, j int) bool {
	return l.Url[i].Priority < l.Url[j].Priority
}

func (l *Sitemap) Swap(i, j int) {
	l.Url[i], l.Url[j] = l.Url[j], l.Url[i]
}
