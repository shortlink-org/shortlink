package uow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	uow "github.com/shortlink-org/shortlink/internal/pkg/uow/mocks"
)

//go:generate mockery

type Entity struct {
	ID   int
	Name string
}

func TestCommit(t *testing.T) {
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}
	ctx := context.Background()

	unitOfWorkMock.On("Commit", ctx).Return(nil)

	err := unitOfWorkMock.Commit(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRollback(t *testing.T) {
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}
	ctx := context.Background()

	unitOfWorkMock.On("Rollback", ctx).Return(nil)

	err := unitOfWorkMock.Rollback(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRegisterNew(t *testing.T) {
	entity := Entity{ID: 1, Name: "Test Entity"}
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}

	unitOfWorkMock.On("RegisterNew", entity).Return(nil)

	err := unitOfWorkMock.RegisterNew(entity)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}
