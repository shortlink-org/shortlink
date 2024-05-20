package v1

func (t *Tariffs) Len() int {
	return len(t.GetList())
}

func (t *Tariffs) Less(i, j int) bool {
	return t.GetList()[i].GetName() < t.GetList()[j].GetName()
}

func (t *Tariffs) Swap(i, j int) {
	t.List[i], t.List[j] = t.GetList()[j], t.GetList()[i]
}
