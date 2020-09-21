package ram

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/batch"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/db/options"
)

// Config ...
type Config struct { // nolint unused
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct { // nolint unused
	// sync.Map solver problem with cache contention
	links sync.Map

	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.setConfig()

	// Create batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			if len(args) == 0 {
				return nil
			}

			for key := range args {
				source := args[key].Item.(*link.Link)
				data, errSingleWrite := s.singleWrite(ctx, source)
				if errSingleWrite != nil {
					return errSingleWrite
				}

				args[key].CB <- data
			}

			return nil
		}

		var err error
		s.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go s.config.job.Run(ctx)
	}

	return nil
}

// Get ...
func (ram *Store) Get(_ context.Context, id string) (*link.Link, error) {
	response, ok := ram.links.Load(id)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	v, ok := response.(*link.Link)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return v, nil
}

// List ...
func (ram *Store) List(_ context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	links := []*link.Link{}

	ram.links.Range(func(key interface{}, value interface{}) bool {
		link, ok := value.(*link.Link)
		if !ok {
			return false
		}

		// Apply Filter
		if isFilterSuccess(link, filter) {
			links = append(links, link)
		}
		return true
	})

	return links, nil
}

// Add ...
func (ram *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch ram.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := ram.config.job.Push(source)
		if err != nil {
			return nil, err
		}

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *link.Link:
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
func (ram *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *Store) Delete(_ context.Context, id string) error {
	ram.links.Delete(id)
	return nil
}

func (ram *Store) singleWrite(_ context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.links.Store(data.Hash, data)

	return data, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
