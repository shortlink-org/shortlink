package uow

import (
	"context"
)

func New[T any](ctx context.Context, unitOfWork UnitOfWork[T]) *Query[T] {
	return &Query[T]{
		UnitOfWork: unitOfWork,
		ctx:        ctx,
		New:        []T{},
		Modified:   []T{},
		Deleted:    []T{},
	}
}

func (q *Query[T]) RegisterNew(entity ...T) error {
	q.New = append(q.New, entity...)
	return nil
}

func (q *Query[T]) RegisterDirty(entity ...T) error {
	q.Modified = append(q.Modified, entity...)
	return nil
}

func (q *Query[T]) RegisterClean(_ ...T) error {
	return nil
}

func (q *Query[T]) RegisterDeleted(entity ...T) error {
	q.Deleted = append(q.Deleted, entity...)
	return nil
}

func (q *Query[T]) Commit() error {
	return q.UnitOfWork.Commit(q.ctx)
}

func (q *Query[T]) Rollback() error {
	return q.UnitOfWork.Rollback(q.ctx)
}
