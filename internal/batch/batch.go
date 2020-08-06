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

	c.mx.Lock()
	c.items = append(c.items, el)
	c.mx.Unlock()

	return el.CB, nil
}

// run - starts a loop flushing at the Interval
func (c *Config) run(ctx context.Context) {
	ticker := time.NewTicker(c.Interval)

	for {
		select {
		case <-ctx.Done():
			for key := range c.items {
				c.items[key].CB <- "ctx close"
			}

			break
		case <-ticker.C:
			c.mx.Lock()

			// skip if items empty
			if len(c.items) > 0 {
				// apply func for all items
				c.cb(c.items)

				// clear items
				c.items = []*Item{}
			}

			c.mx.Unlock()
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
