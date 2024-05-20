package tariff_application

import (
	"context"

	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/tariff/v1"
	tariff_repository "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/repository/tariff"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

type TariffService struct {
	log logger.Logger

	// Subscriber
	// notify.Subscriber[link.Link]

	// Repositories
	tariffRepository tariff_repository.Repository
}

func New(ctx context.Context, log logger.Logger, conn db.DB) (*TariffService, error) {
	// Init Repository
	tariffRepository, err := tariff_repository.New(ctx, conn)
	if err != nil {
		return nil, err
	}

	service := &TariffService{
		log: log,

		// Repositories
		tariffRepository: tariffRepository,
	}

	// Subscribe to Event ==============================================================================================
	notify.Subscribe(uint32(domain.Event_EVENT_TARIFF_NEW), service)
	notify.Subscribe(uint32(domain.Event_EVENT_TARIFF_UPDATE), service)
	notify.Subscribe(uint32(domain.Event_EVENT_TARIFF_CLOSE), service)

	log.Info("init usecase", field.Fields{
		"name": "tariff",
	})

	return service, nil
}

// Notify - implementation of notify.Subscriber interface
func (t *TariffService) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	return notify.Response[any]{}
}

func (t *TariffService) Get(ctx context.Context, id string) (*domain.Tariff, error) {
	return t.tariffRepository.Get(ctx, id)
}

func (t *TariffService) List(ctx context.Context, filter any) (*domain.Tariffs, error) {
	return t.tariffRepository.List(ctx, filter)
}

func (t *TariffService) Add(ctx context.Context, in *domain.Tariff) (*domain.Tariff, error) {
	// generate uniq identity
	in.Id = uuid.New().String()

	return t.tariffRepository.Add(ctx, in)
}

func (t *TariffService) Update(ctx context.Context, in *domain.Tariff) (*domain.Tariff, error) {
	return t.tariffRepository.Update(ctx, in)
}

func (t *TariffService) Delete(ctx context.Context, id string) error {
	return t.tariffRepository.Delete(ctx, id)
}
