package simplecache

import "sync"

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string) bool
}

type cache struct {
	mu    *sync.Mutex
	items map[string]interface{}
}

func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exist := c.items[key]; !exist {
		c.items[key] = value
	}
}

func (c *cache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exist := c.items[key]; exist {
		return item
	}
	return nil
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

func NewCache() Cache {
	return &cache{
		mu:    &sync.Mutex{},
		items: make(map[string]interface{}),
	}
}
