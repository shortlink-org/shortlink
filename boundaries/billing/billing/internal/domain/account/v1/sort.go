package v1

func (a *Accounts) Len() int {
	return len(a.GetList())
}

// Less compares the userId fields of two Account objects and returns true if the userId of the account at index i is less than the userId of the account at index j
func (a *Accounts) Less(i, j int) bool {
	return a.GetList()[i].GetUserId().String() < a.GetList()[j].GetUserId().String()
}

func (a *Accounts) Swap(i, j int) {
	a.list[i], a.list[j] = a.GetList()[j], a.GetList()[i]
}
