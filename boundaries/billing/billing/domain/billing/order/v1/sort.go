package v1

func (o *Orders) Len() int {
	return len(o.GetList())
}

func (o *Orders) Less(i, j int) bool {
	return o.GetList()[i].GetUserId() < o.GetList()[j].GetUserId()
}

func (o *Orders) Swap(i, j int) {
	o.List[i], o.List[j] = o.GetList()[j], o.GetList()[i]
}
