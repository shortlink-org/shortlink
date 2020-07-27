package batch

import (
	"context"
	"time"
)

func New(ctx context.Context, f func(interface{}) interface{}) (*Config, error) {
	cnf := Config{
		ctx:     ctx,
		cb:      f,
		Timeout: time.Second,
	}

	// run background job
	go cnf.run()

	return &cnf, nil
}

func (c *Config) Push(item interface{}) (chan interface{}, error) {
	// create new item
	el := NewItem(item)
	c.items = append(c.items, el)

	//go func(el *Item){
	//	response := c.cb(item)
	//	el.cb <- response
	//}(el)

	return el.cb, nil
}

func (c *Config) run() {
	ticker := time.NewTicker(c.Timeout)

	for {
		select {
		case <-ticker.C:
			c.mx.Lock()
			// apply func for each item
			for _, item := range c.items {
				resp := c.cb(item.item)
				item.cb <- resp
			}

			// clear items
			c.items = []*Item{}
			c.mx.Unlock()
			continue
		case <-c.ctx.Done():
		}
		break
	}
}

func NewItem(item interface{}) *Item {
	return &Item{
		cb:   make(chan interface{}),
		item: item,
	}
}
