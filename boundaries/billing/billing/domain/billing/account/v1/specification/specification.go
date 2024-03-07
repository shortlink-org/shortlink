package specification

import (
	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/domain/billing/account/v1"
)

type Account interface {
	IsSatisfiedBy(account v1.Account) bool
}

type AndAccount struct {
	specs []Account
}

func (a *AndAccount) IsSatisfiedBy(account v1.Account) bool {
	for _, spec := range a.specs {
		if !spec.IsSatisfiedBy(account) {
			return false
		}
	}

	return true
}

func NewAndAccount(specs ...Account) *AndAccount {
	return &AndAccount{
		specs: specs,
	}
}
