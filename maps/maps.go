// Package maps contains generic functions for maps
package maps

// Pair ...
type Pair[K, V any] struct {
	Key   K
	Value V
}

// ToSlice converts a map to a list of
// key/value pairs.
func ToSlice[K, V comparable](m map[K]V) (result []Pair[K, V]) {
	for k, v := range m {
		result = append(result, Pair[K, V]{Key: k, Value: v})
	}
	return result
}
