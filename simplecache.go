package simplecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) bool
}

type data struct {
	value interface{}
	ttl   int64
}

type cache struct {
	mu    *sync.Mutex
	items map[string]*data

	cleaningTicker     *time.Ticker
	cleaningTickerStop chan struct{}
}

func (c *cache) Set(key string, value interface{}, expirationTime time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expired := time.Now().Add(expirationTime).Unix()

	if _, exist := c.items[key]; !exist {
		c.items[key] = &data{
			value: value,
			ttl:   expired,
		}
	}
}

func (c *cache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exist := c.items[key]; exist {
		return item.value, nil
	}
	return nil, fmt.Errorf("do not exist item with key [%s]", key)
}

func (c *cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exist := c.items[key]; exist {
		delete(c.items, key)
		return true
	}

	return false
}

func (c *cache) startCleaner() {
	go func() {
		for {
			select {
			case <-c.cleaningTicker.C:
				c.clean()
			case <-c.cleaningTickerStop:
				return
			}
		}
	}()
}

func (c *cache) clean() {
	now := time.Now().Unix()

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.items {
		if now > value.ttl {
			delete(c.items, key)
		}
	}
}

func NewCache() Cache {
	c := &cache{
		items:              make(map[string]*data),
		mu:                 &sync.Mutex{},
		cleaningTicker:     time.NewTicker(10 * time.Millisecond),
		cleaningTickerStop: make(chan struct{}),
	}
	c.startCleaner()
	return c
}
