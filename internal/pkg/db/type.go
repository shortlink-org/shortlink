package db

import (
	"context"
)

// DB - common interface of db
type DB interface {
	Init(ctx context.Context) error
	GetConn() any
}

// Store abstract type
type Store struct {
	DB

	typeStore string
}
