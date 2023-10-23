package sync_map

import (
	"sync"
)

type SyncMap struct {
	mu sync.RWMutex

	m map[any]any
}

func New() *SyncMap {
	return &SyncMap{
		m: make(map[any]any),
	}
}

func (s *SyncMap) Get(key any) any {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.m[key]
}

func (s *SyncMap) Set(key, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}
