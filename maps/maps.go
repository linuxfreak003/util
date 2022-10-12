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

// Values returns a slice of all the values in a map
func Values[K comparable, V any](m map[K]V) (values []V) {
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
