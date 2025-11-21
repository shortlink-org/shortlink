package ram

import (
	"context"
	"sync"

	"github.com/spf13/viper"

	"github.com/shortlink-org/go-sdk/batch"
	"github.com/shortlink-org/go-sdk/config"
	"github.com/shortlink-org/go-sdk/db/options"

	domain "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/ram/filter"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/types/v1"
)

// Config - config
type Config struct {
	job  *batch.Batch[*domain.Link]
	mode int
}

// Store implementation of db interface
type Store struct {
	config Config
	links  sync.Map
}

// New store
func New(ctx context.Context, cfg *config.Config) (*Store, error) {
	s := &Store{}

	// Set configuration
	s.setConfig()

	// Create a batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(items []*batch.Item[*domain.Link]) error {
			if len(items) == 0 {
				return nil
			}

			for _, item := range items {
				data, errSingleWrite := s.singleWrite(ctx, item.Item)
				if errSingleWrite != nil {
					return errSingleWrite
				}

				item.CallbackChannel <- data
			}

			return nil
		}

		var err error
		s.config.job, err = batch.NewSync[*domain.Link](ctx, cfg, cb)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Get - get
func (s *Store) Get(_ context.Context, id string) (*domain.Link, error) {
	response, ok := s.links.Load(id)
	if !ok {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	v, ok := response.(*domain.Link)
	if !ok {
		return nil, &v1.NotFoundByHashError{Hash: id}
	}

	return v, nil
}

// List - list
func (s *Store) List(_ context.Context, params *v1.FilterLink) (*domain.Links, error) {
	links := domain.NewLinks()

	// Set default filter
	search := filter.NewFilter(params)

	s.links.Range(func(key, value any) bool {
		link, ok := value.(*domain.Link)
		if !ok {
			return false
		}

		// Apply Filter
		if params == nil || search.BuildRAMFilter(link) {
			links.Push(link)
		}

		return true
	})

	return links, nil
}

// Add - add
func (s *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		resCh := s.config.job.Push(source)

		select {
		case res, ok := <-resCh:
			if !ok || res == nil {
				return nil, ErrWrite
			}
			return res, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	case options.MODE_SINGLE_WRITE:
		data, err := s.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Update - update
func (s *Store) Update(_ context.Context, _ *domain.Link) (*domain.Link, error) {
	return nil, nil
}

// Delete - delete
func (s *Store) Delete(_ context.Context, id string) error {
	s.links.Delete(id)

	return nil
}

func (s *Store) singleWrite(_ context.Context, source *domain.Link) (*domain.Link, error) {
	s.links.Store(source.GetHash(), source)

	return source, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // Mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
