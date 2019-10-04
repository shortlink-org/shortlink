package store

import (
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/link"
	"sync"
)

type RamLinkList struct {
	links map[string]link.Link
	mu    sync.Mutex
}

func (l *RamLinkList) Init() error {
	l.mu.Lock()
	l.links = make(map[string]link.Link)
	l.links["test"] = link.Link{Url: "test"}
	l.mu.Unlock()
	return nil
}

func (l RamLinkList) Get(id string) (*link.Link, error) {
	l.mu.Lock()
	response := l.links[id]
	l.mu.Unlock()

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response, nil
}

func (l RamLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.GetHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	l.mu.Lock()
	l.links[data.Hash] = data
	l.mu.Unlock()

	return &data, nil
}

func (l RamLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (l RamLinkList) Delete(id string) error {
	l.mu.Lock()
	delete(l.links, id)
	l.mu.Unlock()

	return nil
}
