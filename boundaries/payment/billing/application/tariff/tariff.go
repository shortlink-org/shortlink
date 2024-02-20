package tariff_application

import (
	"context"

	"github.com/google/uuid"

	billing "github.com/shortlink-org/shortlink/boundaries/payment/billing/domain/billing/tariff/v1"
	tariff_repository "github.com/shortlink-org/shortlink/boundaries/payment/billing/infrastructure/repository/tariff"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

type TariffService struct {
	log logger.Logger

	// Repositories
	tariffRepository tariff_repository.Repository
}

func New(log logger.Logger, tariffRepository tariff_repository.Repository) (*TariffService, error) {
	return &TariffService{
		log: log,

		// Repositories
		tariffRepository: tariffRepository,
	}, nil
}

func (t *TariffService) Get(ctx context.Context, id string) (*billing.Tariff, error) {
	return t.tariffRepository.Get(ctx, id)
}

func (t *TariffService) List(ctx context.Context, filter any) (*billing.Tariffs, error) {
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
