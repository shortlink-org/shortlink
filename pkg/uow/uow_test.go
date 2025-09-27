package uow

import (
	"testing"

	"github.com/stretchr/testify/assert"

	uow "github.com/shortlink-org/shortlink/pkg/uow/mocks"
)

//go:generate mockery

type Entity struct {
	ID   int
	Name string
}

func TestCommit(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "uow")
	t.Attr("component", "uow")

		t.Attr("type", "unit")
		t.Attr("package", "uow")
		t.Attr("component", "uow")
	
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}
	ctx := t.Context()

	unitOfWorkMock.On("Commit", ctx).Return(nil)

	err := unitOfWorkMock.Commit(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRollback(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "uow")
	t.Attr("component", "uow")

		t.Attr("type", "unit")
		t.Attr("package", "uow")
		t.Attr("component", "uow")
	
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}
	ctx := t.Context()

	unitOfWorkMock.On("Rollback", ctx).Return(nil)

	err := unitOfWorkMock.Rollback(ctx)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}

func TestRegisterNew(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "uow")
	t.Attr("component", "uow")

		t.Attr("type", "unit")
		t.Attr("package", "uow")
		t.Attr("component", "uow")
	
	entity := Entity{ID: 1, Name: "Test Entity"}
	unitOfWorkMock := &uow.UnitOfWork[Entity]{}

	unitOfWorkMock.On("RegisterNew", entity).Return(nil)

	err := unitOfWorkMock.RegisterNew(entity)

	assert.NoError(t, err)
	unitOfWorkMock.AssertExpectations(t)
}
