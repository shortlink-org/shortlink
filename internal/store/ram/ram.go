package ram

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/batch"
	"github.com/batazor/shortlink/internal/store/options"
	"github.com/batazor/shortlink/internal/store/query"
)

// RAMConfig ...
type RAMConfig struct { // nolint unused
	mode int
	job  *batch.Config
}

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	// sync.Map solver problem with cache contention
	links sync.Map

	config RAMConfig
}

// Init ...
func (ram *RAMLinkList) Init(ctx context.Context) error { // nolint unparam
	var err error

	// Set configuration
	ram.setConfig()

	// Create batch job
	if ram.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			if len(args) == 0 {
				return nil
			}

			for key := range args {
				source := args[key].Item.(*link.Link)
				data, errSingleWrite := ram.singleWrite(ctx, source)
				if errSingleWrite != nil {
					return errSingleWrite
				}

				args[key].CB <- data
			}

			return nil
		}
		ram.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go ram.config.job.Run(ctx)
	}

	return nil
}

// Close ...
func (ram *RAMLinkList) Close() error {
	return nil
}

// Migrate ...
func (ram *RAMLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (ram *RAMLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
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
func (ram *RAMLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
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
func (ram *RAMLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch ram.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := ram.config.job.Push(source)
		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, err
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
func (ram *RAMLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *RAMLinkList) Delete(ctx context.Context, id string) error { // nolint unused
	ram.links.Delete(id)
	return nil
}

func (ram *RAMLinkList) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.links.Store(data.Hash, data)

	return data, nil
}

// setConfig - set configuration
func (ram *RAMLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_RAM_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to store

	ram.config = RAMConfig{
		mode: viper.GetInt("STORE_RAM_MODE_WRITE"),
	}
}
