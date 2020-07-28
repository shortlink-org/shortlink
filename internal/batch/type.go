package batch

import (
	"sync"
	"time"
)

// Config
type Config struct {
	mx sync.Mutex

	Size     int           // Size is the max entries limit
	Interval time.Duration // Interval is the flush interval
	Worker   int
	Retries  int // Retries is count for try fault exec a command

	items []*Item
	cb    func([]*Item) interface{} // is the flush callback function used to flush entries.
}

// Item
type Item struct {
	CB   chan interface{}
	Item interface{}
}
