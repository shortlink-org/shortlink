package store

import (
	"io"

	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// DB - common interface of store
type DB interface { // nolint unused
	// Closer is the interface that wraps the basic Close method.
	io.Closer

	Init() error

	Get(id string) (*link.Link, error)
	List(filter *query.Filter) ([]*link.Link, error)
	Add(data *link.Link) (*link.Link, error)
	Update(data *link.Link) (*link.Link, error)
	Delete(id string) error
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
	store     DB

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
