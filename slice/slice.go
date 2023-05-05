// Package slice contains generic utility
// functions on slices
package slice

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Max returns the maximum element in a slice
func Max[A constraints.Ordered](as []A) A {
	return MinFunc(as, func(a, b A) bool { return a > b })
}

// Min returns the minimum element in a slice
func Min[A constraints.Ordered](as []A) A {
	return MinFunc(as, func(a, b A) bool { return a < b })
}

// MaxFunc returns the maximum element in a slice
// using a given "less than" function
func MaxFunc[A any](as []A, less func(A, A) bool) A {
	return MinFunc(as, func(a, b A) bool { return less(b, a) })
}

// MinFunc returns the minimum element in a slice
// using a given "less than function"
func MinFunc[A any](as []A, less func(A, A) bool) A {
	var result A
	if len(as) == 0 {
		return result
	}
	result = as[0]
	for _, a := range as {
		if less(a, result) {
			result = a
		}
	}
	return result
}

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
	return mergeSort(l, less)
}

func mergeSort[T any](l []T, less func(T, T) bool) []T {
	// Base case
	if len(l) <= 1 {
		return l
	}

	// Split slice in half
	mid := len(l) / 2
	left := mergeSort(l[:mid], less)
	right := mergeSort(l[mid:], less)

	// Merge contents of each side
	result := make([]T, len(left)+len(right))
	var i, j, r int
	for ; i < len(left) && j < len(right); r++ {
		if less(left[i], right[j]) {
			result[r] = left[i]
			i++
		} else {
			result[r] = right[j]
			j++
		}
	}

	// Finish adding any remaining items
	for ; i < len(left); i++ {
		result[r] = left[i]
		r++
	}
	for ; j < len(right); j++ {
		result[r] = right[j]
		r++
	}
	return result
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

/* TODO: Add options here
Options:
* KeepFirst
* KeepLast int
* MergeFunc func (T, T) T
* HashFunc func (T) []byte
*/

// Deduplicate removes duplicates in a slice
func Deduplicate[T comparable](in []T) []T {
	seen := map[T]struct{}{}
	return Filter(in, func(t T) bool {
		if _, ok := seen[t]; ok {
			return false
		}
		seen[t] = struct{}{}
		return true
	})
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
	return Filter(xs, func(t T) bool {
		return t != v
	})
}

// Reverse reverses the order of a slice
func Reverse[T any](a []T) []T {
	for i, l := 0, len(a); i < l/2; i++ {
		pos := l - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// Intersection returns the intersection of two slices,
// a slice of everything contained in both slices.
func Intersection[T comparable](as []T, bs []T) []T {
	return Filter(as, func(a T) bool {
		return Contains(bs, a)
	})
}

/*
Union returns the union of two slices

Duplicates in the first slice will be left.
Duplicates in the second slice will be removed.
*/
func Union[T comparable](as []T, bs []T) []T {
	for _, b := range bs {
		if !Contains(as, b) {
			as = append(as, b)
		}
	}
	return as
}

// Difference returns the difference of two sets (A - B)
func Difference[T comparable](as []T, bs []T) []T {
	return Filter(as, func(a T) bool {
		return !Contains(bs, a)
	})
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
