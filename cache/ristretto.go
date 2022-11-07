package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

// RistrettoNew returns a new ristretto cache
// which is more performant
func RistrettoNew[K comparable, V any](_ K, _ V, config *ristretto.Config) Cache[K, V] {
	c, err := ristretto.NewCache(config)
	if err != nil {
		panic(err)
	}
	return &ristrettoCache[K, V]{
		cache: c,
	}
}

type ristrettoCache[K comparable, V any] struct {
	cache *ristretto.Cache
}

func (*ristrettoCache[K, V]) Set(_ K, _ V, _ time.Duration) bool { return false }

func (*ristrettoCache[K, V]) Get(_ K) (V, bool) {
	var v V
	return v, false
}
