package sync_map

import (
	"sync"
)

type SyncMap struct {
	mu sync.RWMutex

	m map[interface{}]interface{}
}

func New() *SyncMap {
	return &SyncMap{
		m: make(map[interface{}]interface{}),
	}
}

func (s *SyncMap) Get(key interface{}) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.m[key]
}

func (s *SyncMap) Set(key interface{}, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}
