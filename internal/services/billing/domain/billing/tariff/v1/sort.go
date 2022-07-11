package v1

func (t *Tariffs) Len() int {
	return len(t.List)
}

func (t *Tariffs) Less(i, j int) bool {
	return t.List[i].Name < t.List[j].Name
}

func (t *Tariffs) Swap(i, j int) {
	t.List[i], t.List[j] = t.List[j], t.List[i]
}
