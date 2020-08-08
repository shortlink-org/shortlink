package batch

import (
	"context"
	"time"
)

// TODO: add config for as timeout, retries, etc...
// New - create a new batch
func New(ctx context.Context, cb func([]*Item) interface{}) (*Config, error) {
	cnf := Config{
		cb:       cb,
		Interval: time.Millisecond * 100,
	}

	// run background job
	go cnf.run(ctx)

	return &cnf, nil
}

func (c *Config) Push(item interface{}) (chan interface{}, error) {
	// create new item
	el := NewItem(item)

	c.Lock()
	c.items = append(c.items, el)
	c.Unlock()

	return el.CB, nil
}

// run - starts a loop flushing at the Interval
func (c *Config) run(ctx context.Context) {
	ticker := time.NewTicker(c.Interval)

	for {
		select {
		case <-ctx.Done():
			// skip if items empty
			for key := range c.items {
				c.Lock()
				c.items[key].CB <- "ctx close"
				c.Unlock()
			}
		case <-ticker.C:
			// skip if items empty
			if len(c.items) > 0 {
				c.Lock()
				// apply func for all items
				c.cb(c.items)

				// clear items
				c.items = []*Item{}
				c.Unlock()
			}
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
