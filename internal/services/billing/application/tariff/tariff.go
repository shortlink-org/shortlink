package tariff_application

import (
	"context"

	"github.com/google/uuid"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/tariff/v1"
	billing_store "github.com/shortlink-org/shortlink/internal/services/billing/infrastructure/store"
)

type TariffService struct {
	logger logger.Logger

	// Repositories
	tariffRepository billing_store.TariffRepository
}

func New(logger logger.Logger, tariffRepository billing_store.TariffRepository) (*TariffService, error) {
	return &TariffService{
		logger: logger,

		// Repositories
		tariffRepository: tariffRepository,
	}, nil
}

func (t *TariffService) Get(ctx context.Context, id string) (*billing.Tariff, error) {
	return t.tariffRepository.Get(ctx, id)
}

func (t *TariffService) List(ctx context.Context, filter interface{}) (*billing.Tariffs, error) {
	return t.tariffRepository.List(ctx, filter)
}

func (t *TariffService) Add(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error) {
	// generate uniq identity
	in.Id = uuid.New().String()

	return t.tariffRepository.Add(ctx, in)
}

func (t *TariffService) Update(ctx context.Context, in *billing.Tariff) (*billing.Tariff, error) {
	return t.tariffRepository.Update(ctx, in)
}

func (t *TariffService) Delete(ctx context.Context, id string) error {
	return t.tariffRepository.Delete(ctx, id)
}
