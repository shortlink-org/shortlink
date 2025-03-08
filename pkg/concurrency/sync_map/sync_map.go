package sync_map

import (
	"sync/atomic"
)

type SyncMap struct {
	// We use atomic.Value to store the map, so we don't need to use mutex
	m atomic.Value
}

func New() *SyncMap {
	m := make(map[any]any)
	v := atomic.Value{}
	v.Store(m)

	return &SyncMap{
		m: v,
	}
}

// Get returns the value stored in the map for a key
func (s *SyncMap) Get(key any) any {
	m, ok := s.m.Load().(map[any]any)
	if !ok {
		return nil
	}

	return m[key]
}

// Set stores a value in the map for a key
func (s *SyncMap) Set(key, value any) {
	m, ok := s.m.Load().(map[any]any)
	if !ok {
		return
	}

	m[key] = value
	s.m.Store(m)
}
