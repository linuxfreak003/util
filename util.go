// Package util is a library of generic
// utility functions.
package util

import (
	"math"
)

// Signed is any signed types
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is any insigned integer types
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is any integer types
type Integer interface {
	Signed | Unsigned
}

// Float is any floating point types
type Float interface {
	~float32 | ~float64
}

// Number is any number type
type Number interface {
	Integer | Float
}

// Round rounds a float to the specified digits of precision
func Round[N Number](f float64, p N) float64 {
	x := math.Pow(10, float64(p))
	return math.Round(f*x) / x
}
