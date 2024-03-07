package specification

import (
	"errors"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/domain/billing/account/v1"
)

type Account interface {
	IsSatisfiedBy(account *v1.Account) error
}

type AndAccount struct {
	specs []Account
}

func (a *AndAccount) IsSatisfiedBy(account *v1.Account) error {
	var errs error

	for _, spec := range a.specs {
		if err := spec.IsSatisfiedBy(account); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

func NewAndAccount(specs ...Account) *AndAccount {
	return &AndAccount{
		specs: specs,
	}
}
