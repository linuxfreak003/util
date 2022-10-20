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
	cache map[K]*entry[V]
	mutex *sync.Mutex
}

// entry is the entity saved in the cache
type entry[V any] struct {
	Value       V
	LastUpdated time.Time
	TTL         time.Duration
}

// New returns a new Cache
// It takes any instance of the Key and Value
// to be able to infer the types
func New[K comparable, V any](_ K, _ V) Cache[K, V] {
	impl := &cacheImpl[K, V]{
		cache: make(map[K]*entry[V]),
		mutex: &sync.Mutex{},
	}

	return impl
}

// NewWithGarbageCollection runs a goroutine to clear
// up any stale entries on a given interval
func NewWithGarbageCollection[K comparable, V any](_ K, _ V, d time.Duration) Cache[K, V] {
	impl := &cacheImpl[K, V]{
		cache: make(map[K]*entry[V]),
		mutex: &sync.Mutex{},
	}

	go func() {
		for {
			time.Sleep(d)
			impl.removeStale()
		}
	}()

	return impl
}

func (c *cacheImpl[K, V]) removeStale() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, v := range c.cache {
		if v.TTL > 0 && time.Now().Sub(v.LastUpdated) > v.TTL {
			delete(c.cache, k)
		}
	}
}

// Set sets a value in the cache
func (c *cacheImpl[K, V]) Set(key K, value V, ttl time.Duration) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = &entry[V]{
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
