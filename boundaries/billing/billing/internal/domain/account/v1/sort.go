package v1

func (a *Accounts) Len() int {
	return len(a.GetList())
}

func (a *Accounts) Less(i, j int) bool {
	return a.GetList()[i].GetUserId() < a.GetList()[j].GetUserId()
}

func (a *Accounts) Swap(i, j int) {
	a.list[i], a.list[j] = a.GetList()[j], a.GetList()[i]
}
