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

func (s *SyncMap) Get(key any) any {
	return s.m.Load().(map[any]any)[key]
}

func (s *SyncMap) Set(key, value any) {
	m, ok := s.m.Load().(map[any]any)
	if !ok {
		return
	}

	m[key] = value
	s.m.Store(m)
}
