// Package slice contains generic utility
// functions on slices
package slice

import (
	"github.com/linuxfreak003/util/number"
	"golang.org/x/exp/constraints"
)

// BubbleSort will sort a slice
func BubbleSort[T constraints.Ordered](ts []T) []T {
	for x := 0; x < len(ts); x++ {
		for i := 0; i < len(ts)-1; i++ {
			if ts[i] > ts[i+1] {
				ts[i], ts[i+1] = ts[i+1], ts[i]
			}
		}
	}
	return ts
}

// DumbSort ...
func DumbSort[T constraints.Ordered](ts []T) []T {
	for i := 0; i < len(ts); i++ {
		v := number.Min(ts[i:]...)
		index := Index(ts, v)
		ts[i], ts[index] = v, ts[i]
	}
	return ts
}
