package v1

func (a *Accounts) Len() int {
	return len(a.List)
}

func (a *Accounts) Less(i, j int) bool {
	return a.List[i].UserId < a.List[j].UserId
}

func (a *Accounts) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}
