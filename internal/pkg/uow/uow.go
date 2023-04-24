package uow

import (
	"context"
)

func New[T any](unitOfWork UnitOfWork[T], ctx context.Context) *Query[T] {
	return &Query[T]{
		UnitOfWork: unitOfWork,
		ctx:        ctx,
		New:        []T{},
		Modified:   []T{},
		Deleted:    []T{},
	}
}

func (q *Query[T]) RegisterNew(entity T) error {
	if err := q.UnitOfWork.RegisterNew(entity); err != nil {
		return err
	}
	q.New = append(q.New, entity)
	return nil
}

func (q *Query[T]) RegisterDirty(entity T) error {
	if err := q.UnitOfWork.RegisterDirty(entity); err != nil {
		return err
	}
	q.Modified = append(q.Modified, entity)
	return nil
}

func (q *Query[T]) RegisterClean(entity T) error {
	if err := q.UnitOfWork.RegisterClean(entity); err != nil {
		return err
	}
	return nil
}

func (q *Query[T]) RegisterDeleted(entity T) error {
	if err := q.UnitOfWork.RegisterDeleted(entity); err != nil {
		return err
	}
	q.Deleted = append(q.Deleted, entity)
	return nil
}

func (q *Query[T]) Commit() error {
	return q.UnitOfWork.Commit(q.ctx)
}

func (q *Query[T]) Rollback() error {
	return q.UnitOfWork.Rollback(q.ctx)
}
