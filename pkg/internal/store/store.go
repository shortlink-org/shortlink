package store

import (
	"github.com/batazor/shortlink/pkg/link"
)

// DB - common interface of store
type DB interface { // nolint unused
	Init() error

	Get(id string) (*link.Link, error)
	Add(data link.Link) (*link.Link, error)
	Update(data link.Link) (*link.Link, error)
	Delete(id string) error
}

// Store abstract type
type Store struct{} // nolint unused

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
		store = &RAMLinkList{}
	default:
		store = &RAMLinkList{}
	}

	if err := store.Init(); err != nil {
		panic(err)
	}

	return store
}
