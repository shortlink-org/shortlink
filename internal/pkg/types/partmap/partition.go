package partmap

import (
	"sync"
)

// partition represents a single partition of the PartMap.
type partition struct {
	mu    sync.RWMutex
	store map[string]any
}

func (p *partition) get(key string) (any, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	val, ok := p.store[key]

	return val, ok
}

func (p *partition) set(key string, val any) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.store[key] = val
}

func (p *partition) delete(key string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.store, key)
}

func (p *partition) len() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.store)
}
