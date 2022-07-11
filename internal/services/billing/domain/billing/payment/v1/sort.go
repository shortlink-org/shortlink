package v1

func (p *Payments) Len() int {
	return len(p.List)
}

func (p *Payments) Less(i, j int) bool {
	return p.List[i].Name < p.List[j].Name
}

func (p *Payments) Swap(i, j int) {
	p.List[i], p.List[j] = p.List[j], p.List[i]
}
