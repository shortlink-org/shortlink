package uow

import (
	"context"
)

type UnitOfWork[T any] interface {
	RegisterNew(in T) error
	RegisterDirty(in T) error
	RegisterClean(in T) error
	RegisterDeleted(in T) error

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

//
//type useCase struct{}
//
//func (uc *useCase) Handle(ctx context.Context, uow UnitOfWork) error {
//	if err := uow.Commit(ctx); err != nil {
//		if err := uow.Rollback(ctx); err != nil {
//			return err
//		}
//		return err
//	}
//	return nil
//}
//
//type UseCase[In any] interface {
//	Handle(ctx context.Context, in In) error
//}
