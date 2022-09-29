// Package slice contains generic utility
// functions on slices
package slice

import "math/rand"

// Map create a new slice from an existing one
// using a map function
func Map[A, B any](as []A, f func(A) B) []B {
	bs := make([]B, len(as))
	for i, a := range as {
		bs[i] = f(a)
	}
	return bs
}

// Fold folds the slice into a single value
// using the given function and accumulator
func Fold[A, B any](as []A, acc B, f func(A, B) B) B {
	for _, a := range as {
		acc = f(a, acc)
	}
	return acc
}

// Contains checks for the existence of an
// element in the slice
func Contains[T comparable](l []T, v T) bool {
	for _, t := range l {
		if t == v {
			return true
		}
	}
	return false
}

// Shuffle shuffles a slice using the Fisher-Yates algorithm
// you need to call `rand.Seed` for this to work properly
func Shuffle[T any](l []T) []T {
	n := len(l)
	for i := 0; i < n; i++ {
		r := rand.Intn(n-i) + i
		l[i], l[r] = l[r], l[i]
	}
	return l
}

// Filter filters a slice using the given function
func Filter[T any](in []T, f func(T) bool) (out []T) {
	for _, t := range in {
		if f(t) {
			out = append(out, t)
		}
	}
	return out
}

// Deduplicate deduplicates a slice
func Deduplicate[T comparable](in []T) []T {
	seen := map[T]struct{}{}
	result := []T{}
	for _, el := range in {
		if _, ok := seen[el]; ok {
			continue
		}
		result = append(result, el)
		seen[el] = struct{}{}
	}
	return result
}

// Remove will remove from the slice
// the first match of the given value
func Remove[T comparable](xs []T, d T) []T {
	for i, x := range xs {
		if x == d {
			return append(xs[:i], xs[i+1:]...)
		}
	}
	return xs
}

// RemoveAll removes all occurances of
// the given value in the slice
func RemoveAll[T comparable](xs []T, v T) []T {
	var result []T
	for _, x := range xs {
		if x != v {
			result = append(result, x)
		}
	}
	return result
}
