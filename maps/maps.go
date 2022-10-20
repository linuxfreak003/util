// Package maps contains generic functions for maps
package maps

// ToSlice converts a map to a list of
// key/value pairs.
func ToSlice[K, V comparable](m map[K]V) (result []struct {
	Key   K
	Value V
}) {
	for k, v := range m {
		result = append(result, struct {
			Key   K
			Value V
		}{k, v})
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

// MapValues transforms the values of the map using the given function
func MapValues[K comparable, X, Y any](m map[K]X, f func(X) Y) map[K]Y {
	result := make(map[K]Y)
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}
