package cacherefresh

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	data     T
	err      error
	lock     sync.Mutex
	loadFunc func() (T, error)
}

func New[T any](loadFunc func() (T, error), refreshInterval time.Duration) (*Cache[T], error) {
	c := &Cache[T]{loadFunc: loadFunc}
	c.Refresh()
	if c.err != nil {
		return nil, c.err
	}

	go func() {
		for range time.NewTicker(refreshInterval).C {
			c.Refresh()
		}
	}()

	return c, nil
}

func (c *Cache[T]) Refresh() {
	data, err := c.loadFunc()
	c.lock.Lock()
	c.data = data
	c.err = err
	c.lock.Unlock()
}

func (c *Cache[T]) Get() (T, error) {
	return c.data, c.err
}
