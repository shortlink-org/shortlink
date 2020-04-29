package ram

import (
	"context"
	"fmt"
	"sync"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	ctx context.Context

	// sync.Map solver problem with cache contention
	links sync.Map
}

// Init ...
func (ram *RAMLinkList) Init(ctx context.Context) error { // nolint unparam

	// Set context
	ram.ctx = ctx

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

		links = append(links, link)
		return true
	})

	return links, nil
}

// Add ...
func (ram *RAMLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.links.Store(data.Hash, data)

	return data, nil
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
