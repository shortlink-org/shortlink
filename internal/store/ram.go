package store

import (
	"fmt"
	"sync"

	"github.com/batazor/shortlink/pkg/link"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	links map[string]link.Link
	mu    sync.Mutex
}

// Init ...
func (ram *RAMLinkList) Init() error {
	ram.mu.Lock()
	ram.links = make(map[string]link.Link)
	ram.mu.Unlock()
	return nil
}

// Close ...
func (ram *RAMLinkList) Close() error {
	return nil
}

// Migrate ...
func (ram *RAMLinkList) migrate() error {
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
func (ram *RAMLinkList) List() ([]*link.Link, error) {
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
func (ram *RAMLinkList) Add(data link.Link) (*link.Link, error) {
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
func (ram *RAMLinkList) Delete(id string) error {
	ram.mu.Lock()
	delete(ram.links, id)
	ram.mu.Unlock()

	return nil
}
