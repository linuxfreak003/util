// Package cache is a very simple cache library
// that supports generic types
package cache

import "sync"

// Cache is a simple cache using a map
type Cache[K comparable, V any] struct {
	cache map[K]V
	mutex sync.Mutex
}

// New returns a new Cache
// It takes any instance of the Key and Value
// to be able to infer the types
func New[K comparable, V any](_ K, _ V) *Cache[K, V] {
	return &Cache[K, V]{
		cache: make(map[K]V),
	}
}

// Set sets a value in the cache
func (c *Cache[K, V]) Set(key K, value V) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = value
	return true
}

// Get gets a value from the cache
func (c *Cache[K, V]) Get(key K) (V, bool) {
	v, b := c.cache[key]
	return v, b
}
