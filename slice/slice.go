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

// ToMap converts a slice to a map using
// the given function which takes an element
// and returns a key and value
func ToMap[K comparable, V, A any](as []A, f func(A) (K, V)) map[K]V {
	result := make(map[K]V)
	for _, a := range as {
		k, v := f(a)
		result[k] = v
	}
	return result
}

// Fold folds the slice into a single value
// using the given function and accumulator
func Fold[A, B any](as []A, f func(A, B) B, acc B) B {
	for _, a := range as {
		acc = f(a, acc)
	}
	return acc
}

// Reduce reduces the slice to a single value
// using the given fucntion and accumulator
func Reduce[A, B any](as []A, f func(A, B) B, acc B) B {
	return Fold(as, f, acc)
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

// Sort sorts a slice
func Sort[T any](l []T, less func(T, T) bool) []T {
	return quickMergeSort(l, less)
}

func quickMergeSort[T any](l []T, less func(T, T) bool) []T {
	if len(l) <= 1 {
		return l
	}
	pivot := l[0]
	var smaller []T
	var bigger []T
	for _, t := range l[1:] {
		if less(t, pivot) {
			smaller = append(smaller, t)
		} else {
			bigger = append(bigger, t)
		}
	}

	return append(append(quickMergeSort(smaller, less), pivot), quickMergeSort(bigger, less)...)
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

// Reverse reverses the order of a slice
func Reverse[T any](a []T) []T {
	for i := 0; i < len(a)/2; i++ {
		pos := len(a) - i - 1
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// Intersect returns the intersection of two slices
func Intersect[T comparable](as []T, bs []T) (ts []T) {
	for _, a := range as {
		if Contains(bs, a) {
			ts = append(ts, a)
		}
	}
	return ts
}

// Index returns the index of a value if it exists
// it will return -1 if it does not exist in the slice
func Index[T comparable](ts []T, v T) int {
	for i, t := range ts {
		if t == v {
			return i
		}
	}
	return -1
}

// MapFilter combines Map and Filter into one process
// It takes a function that takes a value of type A
// and returns a value of type B, and a bool for whether
// or not to keep the value.
func MapFilter[A, B any](slice []A, f func(A) (B, bool)) (out []B) {
	for _, a := range slice {
		if b, keep := f(a); keep {
			out = append(out, b)
		}
	}
	return out
}
