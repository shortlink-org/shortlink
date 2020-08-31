package db

import (
	"context"
	"io"
)

// DB - common interface of db
type DB interface { // nolint unused
	// Closer is the interface that wraps the basic Close method.
	io.Closer

	Init(ctx context.Context) error
	GetConn() interface{}
}

// Store abstract type
type Store struct { // nolint unused
	Store DB

	typeStore string
}
