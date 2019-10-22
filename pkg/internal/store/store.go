package store

import (
	"github.com/batazor/shortlink/pkg/link"
)

type DB interface {
	Init() error

	Get(id string) (*link.Link, error)
	Add(data link.Link) (*link.Link, error)
	Update(data link.Link) (*link.Link, error)
	Delete(id string) error
}

// Store abstract type
type Store struct{}

// Use return implementation of store
func (s *Store) Use() DB {
	var store DB

	typeStore := "ram"

	switch typeStore {
	case "postgres":
		store = &PostgresLinkList{}
	case "mongo":
		store = &MongoLinkList{}
	case "redis":
		store = &RedisLinkList{}
	case "dgraph":
		store = &DGraphLinkList{}
	case "leveldb":
		store = &LevelDBLinkList{}
	case "badger":
		store = &BadgerLinkList{}
	case "ram":
		store = &RamLinkList{}
	default:
		store = &RamLinkList{}
	}

	if err := store.Init(); err != nil {
		panic(err)
	}

	return store
}
