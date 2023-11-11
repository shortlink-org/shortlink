package ram

import (
	"context"
	"sync"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// Config - config
type Config struct {
	job  *batch.Batch
	mode int
}

// Store implementation of db interface
type Store struct {
	config Config
	links  sync.Map
}

// New store
func New(ctx context.Context) (*Store, error) {
	s := &Store{}

	// Set configuration
	s.setConfig()

	// Create a batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) any {
			if len(args) == 0 {
				return nil
			}

			for key := range args {
				source := args[key].Item.(*domain.Link) //nolint:errcheck // ignore
				data, errSingleWrite := s.singleWrite(ctx, source)
				if errSingleWrite != nil {
					return errSingleWrite
				}

				args[key].CallbackChannel <- data
			}

			return nil
		}

		var err error
		s.config.job, err = batch.New(ctx, cb)
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
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}}
	}

	v, ok := response.(*domain.Link)
	if !ok {
		return nil, &domain.NotFoundError{Link: &domain.Link{Hash: id}}
	}

	return v, nil
}

// List - list
func (s *Store) List(_ context.Context, filter *domain.FilterLink) (*domain.Links, error) {
	links := &domain.Links{
		Link: []*domain.Link{},
	}

	s.links.Range(func(key, value any) bool {
		link, ok := value.(*domain.Link)
		if !ok {
			return false
		}

		// Apply Filter
		if isFilterSuccess(link, filter) {
			links.Link = append(links.GetLink(), link)
		}

		return true
	})

	return links, nil
}

// Add - add
func (s *Store) Add(ctx context.Context, source *domain.Link) (*domain.Link, error) {
	switch s.config.mode {
	case options.MODE_BATCH_WRITE:
		cb := s.config.job.Push(source)

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *domain.Link:
			return data, nil
		default:
			return nil, nil
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

// Close - close
func (s *Store) Close() error {
	if s.config.job != nil {
		s.config.job.Stop()
	}

	return nil
}

func (s *Store) singleWrite(_ context.Context, source *domain.Link) (*domain.Link, error) {
	err := domain.NewURL(source) // Create a new link
	if err != nil {
		return nil, err
	}

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
