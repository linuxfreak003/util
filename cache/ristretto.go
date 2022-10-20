package cache

import "github.com/dgraph-io/ristretto"

// RistrettoNew returns a new ristretto cache
// which is more performant
func RistrettoNew[K comparable, V any](_ K, _ V, config *ristretto.Config) Cache[K, V] {
	return *ristrettoCache{
		cache: ristretto.NewCache(config),
	}
}

type ristrettoCache[K,V] struct {
	cache *ristretto.Cache
}

func (c *ristrettoCache[K,V]) Set()
func (c *ristrettoCache) Get()
