package v1

func (o *Orders) Len() int {
	return len(o.GetList())
}

func (o *Orders) Less(i, j int) bool {
	return o.GetList()[i].GetUserId().String() < o.GetList()[j].GetUserId().String()
}

func (o *Orders) Swap(i, j int) {
	o.list[i], o.list[j] = o.GetList()[j], o.GetList()[i]
}
