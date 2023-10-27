package db

import (
	"context"
	"io"
)

// DB - common interface of db
type DB interface {
	// Closer is the interface that wraps the basic Close method.
	io.Closer

	Init(ctx context.Context) error
	GetConn() any
}

// Store abstract type
type Store struct {
	DB

	typeStore string
}
