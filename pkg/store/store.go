package store

import "github.com/batazor/shortlink/pkg/link"

type DB interface {
	Init() error

	Get(id string) (*link.Link, error)
	Add(data link.Link) (*link.Link, error)
	Update(data link.Link) (*link.Link, error)
	Delete(id string) error
}

type Store struct{}

func (s Store) Use() DB {
	return &RamLinkList{}
}
