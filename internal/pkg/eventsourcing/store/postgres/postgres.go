package postgres

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	eventsourcing "github.com/batazor/shortlink/internal/pkg/eventsourcing/v1"
)

type Store struct {
	db *db.Store
}

func (s *Store) Init(ctx context.Context, db *db.Store) (*Store, error) {
	return &Store{
		db: db,
	}, nil
}

func (s *Store) save(events []eventsourcing.Event, version int, safe bool) error {
	if len(events) == 0 {
		return nil
	}

	// Build all event records, with incrementing versions starting from the
	// original aggregate version.

	panic("implement me")
}

func (s *Store) Save(ctx context.Context, events []eventsourcing.Event, version int) error {
	panic("implement me")
}

func (s *Store) SafeSave(ctx context.Context, events []eventsourcing.Event, version int) error {
	panic("implement me")
}

func (s *Store) Load(ctx context.Context, aggregateID string) ([]eventsourcing.Event, error) {
	panic("implement me")
}
