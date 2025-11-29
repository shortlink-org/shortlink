//go:build unit

package link

import (
	"context"
	"errors"
	"testing"

	"github.com/authzed/authzed-go/v1"
	"github.com/shortlink-org/go-sdk/auth/session"
	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/kratos"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud"
)

// MockKratosClient is a mock implementation of kratos.KratosClient interface
type MockKratosClient struct {
	mock.Mock
}

func (m *MockKratosClient) GetUserEmail(ctx context.Context, userID string) (string, error) {
	args := m.Called(ctx, userID)
	return args.String(0), args.Error(1)
}

// Ensure MockKratosClient implements kratos.KratosClient interface
var _ kratos.KratosClient = (*MockKratosClient)(nil)

// MockRepository is a mock implementation of crud.Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Get(ctx context.Context, id string) (*domain.Link, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Link), args.Error(1)
}

func (m *MockRepository) List(ctx context.Context, filter *domain.FilterLink) (*domain.Links, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Links), args.Error(1)
}

func (m *MockRepository) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Link), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Link), args.Error(1)
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Ensure MockRepository implements crud.Repository interface
var _ crud.Repository = (*MockRepository)(nil)

func TestUC_Get(t *testing.T) {
	ctx := context.Background()

	cfg, err := config.New()
	require.NoError(t, err)

	log, _, err := logger.NewDefault(ctx, cfg)
	require.NoError(t, err)

	mockStore := new(MockRepository)
	mockKratos := new(MockKratosClient)
	mockPermission := &authzed.Client{}

	uc := &UC{
		log:        log,
		store:      mockStore,
		kratos:     mockKratos,
		permission: mockPermission,
		eventBus:   nil, // EventBus not needed for Get use case tests
	}

	t.Run("Public link - success", func(t *testing.T) {
		publicLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/public").
			SetDescribe("Public link").
			Build()
		require.NoError(t, err)
		hash := publicLink.GetHash()

		ctxWithUser := session.WithUserID(ctx, "user-123")

		mockStore.On("Get", mock.Anything, hash).Return(publicLink, nil).Once()

		result, err := uc.Get(ctxWithUser, hash)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, hash, result.GetHash())
		assert.True(t, result.IsPublic())
		mockStore.AssertExpectations(t)
	})

	t.Run("Public link - not found", func(t *testing.T) {
		hash := "non-existent-hash"

		ctxWithUser := session.WithUserID(ctx, "user-123")

		mockStore.On("Get", mock.Anything, hash).Return(nil, domain.ErrNotFound(hash)).Once()

		result, err := uc.Get(ctxWithUser, hash)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, errors.Is(err, domain.ErrNotFound("")))
		mockStore.AssertExpectations(t)
	})

	t.Run("Private link - anonymous user - permission denied", func(t *testing.T) {
		privateLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/private").
			SetDescribe("Private link").
			SetAllowedEmails([]string{"allowed@example.com"}).
			Build()
		require.NoError(t, err)
		hash := privateLink.GetHash()

		ctxWithAnonymous := session.WithUserID(ctx, "anonymous")

		mockStore.On("Get", mock.Anything, hash).Return(privateLink, nil).Once()

		result, err := uc.Get(ctxWithAnonymous, hash)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, errors.Is(err, domain.ErrPermissionDenied(nil)))
		mockStore.AssertExpectations(t)
		mockKratos.AssertNotCalled(t, "GetUserEmail")
	})

	t.Run("Private link - authenticated user - email in allowlist - success", func(t *testing.T) {
		privateLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/private-allowed").
			SetDescribe("Private link").
			SetAllowedEmails([]string{"user@example.com"}).
			Build()
		require.NoError(t, err)
		hash := privateLink.GetHash()

		userID := "user-123"
		userEmail := "user@example.com"

		ctxWithUser := session.WithUserID(ctx, userID)

		mockStore.On("Get", mock.Anything, hash).Return(privateLink, nil).Once()
		mockKratos.On("GetUserEmail", mock.Anything, userID).Return(userEmail, nil).Once()

		result, err := uc.Get(ctxWithUser, hash)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, hash, result.GetHash())
		assert.False(t, result.IsPublic())
		mockStore.AssertExpectations(t)
		mockKratos.AssertExpectations(t)
	})

	t.Run("Private link - authenticated user - email not in allowlist - permission denied", func(t *testing.T) {
		privateLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/private-not-allowed").
			SetDescribe("Private link").
			SetAllowedEmails([]string{"allowed@example.com"}).
			Build()
		require.NoError(t, err)
		hash := privateLink.GetHash()

		userID := "user-123"
		userEmail := "not-allowed@example.com"

		ctxWithUser := session.WithUserID(ctx, userID)

		mockStore.On("Get", mock.Anything, hash).Return(privateLink, nil).Once()
		mockKratos.On("GetUserEmail", mock.Anything, userID).Return(userEmail, nil).Once()

		result, err := uc.Get(ctxWithUser, hash)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, errors.Is(err, domain.ErrPermissionDenied(nil)))
		mockStore.AssertExpectations(t)
		mockKratos.AssertExpectations(t)
	})

	t.Run("Private link - Kratos error - permission denied", func(t *testing.T) {
		privateLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/private-kratos-error").
			SetDescribe("Private link").
			SetAllowedEmails([]string{"allowed@example.com"}).
			Build()
		require.NoError(t, err)
		hash := privateLink.GetHash()

		userID := "user-123"
		kratosError := errors.New("failed to get user identity")

		ctxWithUser := session.WithUserID(ctx, userID)

		mockStore.On("Get", mock.Anything, hash).Return(privateLink, nil).Once()
		mockKratos.On("GetUserEmail", mock.Anything, userID).Return("", kratosError).Once()

		result, err := uc.Get(ctxWithUser, hash)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, errors.Is(err, domain.ErrPermissionDenied(nil)))
		mockStore.AssertExpectations(t)
		mockKratos.AssertExpectations(t)
	})

	t.Run("Private link - empty user ID - permission denied", func(t *testing.T) {
		privateLink, err := domain.NewLinkBuilder().
			SetURL("https://example.com/private-empty-user").
			SetDescribe("Private link").
			SetAllowedEmails([]string{"allowed@example.com"}).
			Build()
		require.NoError(t, err)
		hash := privateLink.GetHash()

		ctxWithEmptyUser := session.WithUserID(ctx, "")

		mockStore.On("Get", mock.Anything, hash).Return(privateLink, nil).Once()

		result, err := uc.Get(ctxWithEmptyUser, hash)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, errors.Is(err, domain.ErrPermissionDenied(nil)))
		mockStore.AssertExpectations(t)
		mockKratos.AssertNotCalled(t, "GetUserEmail")
	})
}
