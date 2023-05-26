/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"time"
)

// New creates a new batch Config with a specified callback function.
func New(ctx context.Context, cb func([]*Item) interface{}) (*Config, error) {
	c := &Config{
		ctx:      ctx,
		callback: cb,
		interval: time.Millisecond * 100, // nolint:gomnd
	}

	go c.run()
	return c, nil
}

// Push adds an item to the batch.
func (c *Config) Push(item interface{}) chan interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	newItem := &Item{
		CallbackChannel: make(chan interface{}),
		Item:            item,
	}
	c.items = append(c.items, newItem)
	return newItem.CallbackChannel
}

// run starts a loop flushing at the specified interval.
func (c *Config) run() {
	ticker := time.NewTicker(c.interval)

	for {
		select {
		case <-c.ctx.Done():
			c.mu.Lock()
			for _, item := range c.items {
				item.CallbackChannel <- struct{}{}
			}
			c.mu.Unlock()

		case <-ticker.C:
			c.mu.Lock()
			if len(c.items) > 0 {
				c.callback(c.items)
				c.items = []*Item{}
			}
			c.mu.Unlock()
		}
	}
}
