package cache

import "sync"

type Cache[K comparable, V any] struct {
	cache map[K]V
	mutex sync.Mutex
}

func New[K comparable, V any](_ K, _ V) *Cache[K, V] {
	return &Cache[K, V]{
		cache: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
	return true
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, b := c.cache[key]
	return v, b
}
