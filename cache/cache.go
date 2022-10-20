// Package cache is a very simple cache library
// that supports generic types
package cache

import (
	"sync"
	"time"
)

// Cache is a simple cache interface
type Cache[K comparable, V any] interface {
	Set(K, V, time.Duration) bool
	Get(K) (V, bool)
}

// cacheImpl is a simple cache implementation using a map
type cacheImpl[K comparable, V any] struct {
	cache map[K]*Entry[V]
	mutex *sync.Mutex
}

// Entry is the entity saved in the cache
type Entry[V any] struct {
	Value       V
	LastUpdated time.Time
	TTL         time.Duration
}

// New returns a new Cache
// It takes any instance of the Key and Value
// to be able to infer the types
func New[K comparable, V any](_ K, _ V) Cache[K, V] {
	return &cacheImpl[K, V]{
		cache: make(map[K]*Entry[V]),
		mutex: &sync.Mutex{},
	}
}

// Set sets a value in the cache
func (c *cacheImpl[K, V]) Set(key K, value V, ttl time.Duration) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = &Entry[V]{
		Value:       value,
		LastUpdated: time.Now(),
		TTL:         ttl,
	}
	return true
}

// Get gets a value from the cache
func (c *cacheImpl[K, V]) Get(key K) (V, bool) {
	entry, inCache := c.cache[key]

	var v V
	if !inCache {
		return v, false
	}

	if entry.TTL != 0 && time.Now().Sub(entry.LastUpdated) > entry.TTL {
		return v, false
	}

	return entry.Value, true
}
