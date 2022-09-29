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

// Contains checks for the existence of an
// element in a slice
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

// Filter filters a slice using a given function
func Filter[T any](in []T, f func(T) bool) (out []T) {
	for _, t := range in {
		if f(t) {
			out = append(out, t)
		}
	}
	return out
}
