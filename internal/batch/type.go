package batch

import (
	"context"
	"sync"
	"time"
)

type Config struct {
	ctx context.Context
	mx  sync.Mutex

	Size    int
	Timeout time.Duration
	Worker  int

	items []*Item
	cb    func(interface{}) interface{}
}

type Item struct {
	cb   chan interface{}
	item interface{}
}
