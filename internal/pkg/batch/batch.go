/*
Package for work in batch mode
*/
package batch

import (
	"context"
	"time"
)

// TODO: add config for as timeout, retries, etc...
// New - create a new batch
func New(_ context.Context, cb func([]*Item) interface{}) (*Config, error) {
	cnf := Config{
		cb:       cb,
		Interval: time.Millisecond * 100, // nolint:gomnd
	}

	return &cnf, nil
}

func (c *Config) Push(item interface{}) (chan interface{}, error) {
	// create new item
	el := NewItem(item)

	c.mu.Lock()
	c.items = append(c.items, el)
	c.mu.Unlock()

	return el.CB, nil
}

// run - starts a loop flushing at the Interval
func (c *Config) Run(ctx context.Context) {
	ticker := time.NewTicker(c.Interval)

	for {
		select {
		case <-ctx.Done():
			c.mu.Lock()

			// skip if items empty
			for key := range c.items {
				c.items[key].CB <- "ctx close"
			}

			c.mu.Unlock()
		case <-ticker.C:
			c.mu.Lock()

			// skip if items empty
			if len(c.items) > 0 {
				// apply func for all items
				c.cb(c.items)

				// clear items
				c.items = []*Item{}
			}

			c.mu.Unlock()
		}
	}
}

func NewItem(item interface{}) *Item {
	cb := make(chan interface{})

	return &Item{
		CB:   cb,
		Item: item,
	}
}
