package store

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// DB - common interface of store
type DB interface { // nolint unused
	// Closer is the interface that wraps the basic Close method.
	io.Closer

	Init(ctx context.Context) error

	Get(ctx context.Context, id string) (*link.Link, error)
	List(ctx context.Context, filter *query.Filter) ([]*link.Link, error)
	Add(ctx context.Context, data *link.Link) (*link.Link, error)
	Update(ctx context.Context, data *link.Link) (*link.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
	store     DB

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
