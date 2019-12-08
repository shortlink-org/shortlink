package ram

import (
	"fmt"
	"sync"

	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/query"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	links map[string]link.Link
	mu    sync.Mutex
}

// Init ...
func (ram *RAMLinkList) Init() error { // nolint unparam
	ram.mu.Lock()
	ram.links = make(map[string]link.Link)
	ram.mu.Unlock()

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, ram)
	notify.Subscribe(api_type.METHOD_GET, ram)
	notify.Subscribe(api_type.METHOD_LIST, ram)
	notify.Subscribe(api_type.METHOD_UPDATE, ram)
	notify.Subscribe(api_type.METHOD_DELETE, ram)

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
func (ram *RAMLinkList) Get(id string) (*link.Link, error) {
	ram.mu.Lock()
	response := ram.links[id]
	ram.mu.Unlock()

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (ram *RAMLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	links := []*link.Link{}

	ram.mu.Lock()
	// copy map by assigning elements to new map
	for key := range ram.links {
		links = append(links, &link.Link{
			Url:      ram.links[key].Url,
			Hash:     ram.links[key].Hash,
			Describe: ram.links[key].Describe,
		})
	}
	ram.mu.Unlock()

	return links, nil
}

// Add ...
func (ram *RAMLinkList) Add(data link.Link) (*link.Link, error) { // nolint unused
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	ram.mu.Lock()
	ram.links[data.Hash] = data
	ram.mu.Unlock()

	return &data, nil
}

// Update ...
func (ram *RAMLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *RAMLinkList) Delete(id string) error { // nolint unused
	ram.mu.Lock()
	delete(ram.links, id)
	ram.mu.Unlock()

	return nil
}

// Notify ...
func (ram *RAMLinkList) Notify(event int, payload interface{}) *notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		payload, err := ram.Add(payload.(link.Link))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_GET:
		payload, err := ram.Get(payload.(string))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_LIST:
		payload, err := ram.List(nil)
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_UPDATE:
		payload, err := ram.Update(payload.(link.Link))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_DELETE:
		err := ram.Delete(payload.(string))
		return &notify.Response{
			Payload: nil,
			Error:   err,
		}
	}

	return nil
}
