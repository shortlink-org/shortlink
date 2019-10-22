package store

import (
	"fmt"
	"github.com/batazor/shortlink/pkg/link"
	"sync"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct {
	links map[string]link.Link
	mu    sync.Mutex
}

// Init ...
func (l *RAMLinkList) Init() error {
	l.mu.Lock()
	l.links = make(map[string]link.Link)
	l.mu.Unlock()
	return nil
}

// Get ...
func (l *RAMLinkList) Get(id string) (*link.Link, error) {
	l.mu.Lock()
	response := l.links[id]
	l.mu.Unlock()

	if response.URL == "" {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// Add ...
func (l *RAMLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	l.mu.Lock()
	l.links[data.Hash] = data
	l.mu.Unlock()

	return &data, nil
}

// Update ...
func (l *RAMLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (l *RAMLinkList) Delete(id string) error {
	l.mu.Lock()
	delete(l.links, id)
	l.mu.Unlock()

	return nil
}
