package v1

func (o *Orders) Len() int {
	return len(o.List)
}

func (o *Orders) Less(i, j int) bool {
	return o.List[i].UserId < o.List[j].UserId
}

func (o *Orders) Swap(i, j int) {
	o.List[i], o.List[j] = o.List[j], o.List[i]
}
