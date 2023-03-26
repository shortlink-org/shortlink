package batch

import (
	"sync"
	"time"
)

// Config
type Config struct {
	cb       func([]*Item) interface{}
	items    []*Item
	Size     int
	Interval time.Duration
	Worker   int
	Retries  int
	sync.Mutex
}

// Item
type Item struct {
	CB   chan interface{}
	Item interface{}
}
