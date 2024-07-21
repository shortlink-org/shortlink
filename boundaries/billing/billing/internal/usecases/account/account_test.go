package account_application

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/account/v1"
	account_repository "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/account/mocks"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/config"
)

//go:generate mockery

func TestAccountService_Add(t *testing.T) {
	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create a new account
	account, err := v1.NewAccountBuilder().
		SetUserId(mustNewV7(t)).
		SetTariffId(mustNewV7(t)).
		Build()

	require.NoError(t, err, "Error create a new account")

	// Create a mock repository
	mockRepo := account_repository.NewRepository(t)
	mockRepo.On("Add", mock.Anything, account).Return(account, nil)

	// Create the service with the mock repository
	service := &AccountService{
		log:               log,
		accountRepository: mockRepo,
	}

	// Call the Add method
	result, err := service.Add(context.Background(), account)

	// Assert that the returned account is the one we expected
	assert.NoError(t, err)
	assert.Equal(t, account, result)

	// Assert that the Add method was called with the correct account
	mockRepo.AssertCalled(t, "Add", mock.Anything, account)
}

func TestAccountService_Get(t *testing.T) {
	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create a new account
	account, err := v1.NewAccountBuilder().
		SetUserId(mustNewV7(t)).
		SetTariffId(mustNewV7(t)).
		Build()

	require.NoError(t, err, "Error create a new account")

	// Create a mock repository
	mockRepo := account_repository.NewRepository(t)
	mockRepo.On("Get", mock.Anything, account.GetId()).Return(account, nil)

	// Create the service with the mock repository
	service := &AccountService{
		log:               log,
		accountRepository: mockRepo,
	}

	// Call the Get method
	result, err := service.Get(context.Background(), account.GetId().String())

	// Assert that the returned account is the one we expected
	assert.NoError(t, err)
	assert.Equal(t, account, result)

	// Assert that the Get method was called with the correct account id
	mockRepo.AssertCalled(t, "Get", mock.Anything, account.GetId())
}

func TestAccountService_List(t *testing.T) {
	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create a new account
	account, err := v1.NewAccountBuilder().
		SetUserId(mustNewV7(t)).
		SetTariffId(mustNewV7(t)).
		Build()

	require.NoError(t, err, "Error create a new account")

	// Create a mock repository
	mockRepo := account_repository.NewRepository(t)
	mockRepo.On("list", mock.Anything, mock.Anything).Return([]*v1.Account{account}, nil)

	// Create the service with the mock repository
	service := &AccountService{
		log:               log,
		accountRepository: mockRepo,
	}

	// Call the list method
	result, err := service.List(context.Background(), nil)

	// Assert that the returned account list contains the account we expected
	assert.NoError(t, err)
	assert.Equal(t, []*v1.Account{account}, result)

	// Assert that the list method was called
	mockRepo.AssertCalled(t, "list", mock.Anything, mock.Anything)
}

func TestAccountService_Update(t *testing.T) {
	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create a new account
	account, err := v1.NewAccountBuilder().
		SetUserId(mustNewV7(t)).
		SetTariffId(mustNewV7(t)).
		Build()

	require.NoError(t, err, "Error create a new account")

	// Create a mock repository
	mockRepo := account_repository.NewRepository(t)
	mockRepo.On("Update", mock.Anything, account).Return(account, nil)

	// Create the service with the mock repository
	service := &AccountService{
		log:               log,
		accountRepository: mockRepo,
	}

	// Call the Update method
	result, err := service.Update(context.Background(), account)

	// Assert that the returned account is the one we expected
	assert.NoError(t, err)
	assert.Equal(t, account, result)

	// Assert that the Update method was called with the correct account
	mockRepo.AssertCalled(t, "Update", mock.Anything, account)
}

func TestAccountService_Delete(t *testing.T) {
	// Init logger
	conf := config.Configuration{}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Create a new account
	account, err := v1.NewAccountBuilder().
		SetUserId(mustNewV7(t)).
		SetTariffId(mustNewV7(t)).
		Build()

	require.NoError(t, err, "Error create a new account")

	// Create a mock repository
	mockRepo := account_repository.NewRepository(t)
	mockRepo.On("Delete", mock.Anything, account.GetId()).Return(nil)

	// Create the service with the mock repository
	service := &AccountService{
		log:               log,
		accountRepository: mockRepo,
	}

	// Call the Delete method
	err = service.Delete(context.Background(), account.GetId().String())
	assert.NoError(t, err)

	// Assert that the Delete method was called with the correct account id
	mockRepo.AssertCalled(t, "Delete", mock.Anything, account.GetId())
}

func mustNewV7(t *testing.T) uuid.UUID {
	t.Helper()

	id, err := uuid.NewV7()
	require.NoError(t, err)

	return id
}
