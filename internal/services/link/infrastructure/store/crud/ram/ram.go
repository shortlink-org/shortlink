package ram

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/batch"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Config ...
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
		cb := func(args []*batch.Item) interface{} {
			if len(args) == 0 {
				return nil
			}

			for key := range args {
				source := args[key].Item.(*v1.Link) // nolint:errcheck
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

// Get ...
func (ram *Store) Get(_ context.Context, id string) (*v1.Link, error) {
	response, ok := ram.links.Load(id)
	if !ok {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	v, ok := response.(*v1.Link)
	if !ok {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return v, nil
}

// List ...
func (ram *Store) List(_ context.Context, filter *query.Filter) (*v1.Links, error) {
	links := &v1.Links{
		Link: []*v1.Link{},
	}

	ram.links.Range(func(key, value interface{}) bool {
		link, ok := value.(*v1.Link)
		if !ok {
			return false
		}

		// Apply Filter
		if isFilterSuccess(link, filter) {
			links.Link = append(links.Link, link)
		}

		return true
	})

	return links, nil
}

// Add ...
func (ram *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	switch ram.config.mode {
	case options.MODE_BATCH_WRITE:
		cb := ram.config.job.Push(source)

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *v1.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := ram.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Update ...
func (ram *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *Store) Delete(_ context.Context, id string) error {
	ram.links.Delete(id)
	return nil
}

// Close ...
func (ram *Store) Close() error {
	ram.config.job.Stop()
	return nil
}

func (ram *Store) singleWrite(_ context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.links.Store(source.Hash, source)

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
