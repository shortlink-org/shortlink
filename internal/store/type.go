package store

import (
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// DB - common interface of store
type DB interface { // nolint unused
	notify.Subscriber // Observer interface

	Init() error
	Close() error

	Get(id string) (*link.Link, error)
	List(filter *query.Filter) ([]*link.Link, error)
	Add(data link.Link) (*link.Link, error)
	Update(data link.Link) (*link.Link, error)
	Delete(id string) error
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
}
