package ram

import (
	"fmt"
	"sync"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	// sync.Map solver problem with cache contention
	links sync.Map
}

// Init ...
func (ram *RAMLinkList) Init() error { // nolint unparam
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
func (ram *RAMLinkList) Get(hash string) (*link.Link, error) {
	response, ok := ram.links.Load(hash)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: hash}, Err: fmt.Errorf("Not found id: %s", hash)}
	}

	v, ok := response.(*link.Link)
	if !ok {
		return nil, &link.NotFoundError{Link: &link.Link{Url: hash}, Err: fmt.Errorf("Not found id: %s", hash)}
	}

	return v, nil
}

// List ...
func (ram *RAMLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
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
func (ram *RAMLinkList) Add(source *link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.links.Store(data.Hash, data)

	return data, nil
}

// Update ...
func (ram *RAMLinkList) Update(data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *RAMLinkList) Delete(hash string) error { // nolint unused
	ram.links.Delete(hash)
	return nil
}
