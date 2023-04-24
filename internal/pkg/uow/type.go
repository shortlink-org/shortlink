package uow

import (
	"context"
)

type UnitOfWork[T any] interface {
	RegisterNew(in ...T) error
	RegisterDirty(in ...T) error
	RegisterClean(in ...T) error
	RegisterDeleted(in ...T) error

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Query[T any] struct {
	UnitOfWork[T]

	ctx context.Context

	New      []T
	Modified []T
	Deleted  []T
}
