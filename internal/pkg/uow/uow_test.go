package uow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shortlink-org/shortlink/internal/pkg/uow/mocks"
)

//go:generate mockery --name=UnitOfWork --dir=./ --output=./mocks --outpkg=mocks --exported --case=underscore

type Entity struct {
	ID   int
	Name string
}

func TestCommit(t *testing.T) {
	unitOfWorkMock := &mocks.UnitOfWork[Entity]{}
	ctx := context.Background()

	unitOfWorkMock.On("Commit", ctx).Return(nil)

	err := unitOfWorkMock.Commit(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRollback(t *testing.T) {
	unitOfWorkMock := &mocks.UnitOfWork[Entity]{}
	ctx := context.Background()

	unitOfWorkMock.On("Rollback", ctx).Return(nil)

	err := unitOfWorkMock.Rollback(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRegisterNew(t *testing.T) {
	entity := Entity{ID: 1, Name: "Test Entity"}
	unitOfWorkMock := &mocks.UnitOfWork[Entity]{}

	unitOfWorkMock.On("RegisterNew", entity).Return(nil)

	err := unitOfWorkMock.RegisterNew(entity)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}
