// Package maps contains generic functions for maps
package maps

// ToSlice converts a map to a slice
// using a given function
func ToSlice[K comparable, V, T any](m map[K]V, f func(K, V) T) []T {
	out := make([]T, len(m))
	i := 0
	for k, v := range m {
		out[i] = f(k, v)
		i++
	}
	return out
}

// Values returns a slice of all the values in a map
func Values[K comparable, V any](m map[K]V) []V {
	return ToSlice(m, func(_ K, v V) V {
		return v
	})
}

// Keys returns a slice of all the keys in a map
func Keys[K comparable, V any](m map[K]V) []K {
	return ToSlice(m, func(k K, _ V) V {
		return k
	})
}

// MapValues transforms the values of the map using the given function
func MapValues[K comparable, X, Y any](m map[K]X, f func(X) Y) map[K]Y {
	result := make(map[K]Y)
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}
