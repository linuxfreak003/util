// Package util is a library of generic
// utility functions.
package util

import (
	"math"
)

// Round rounds a float to the specified digits of precision
func Round[N Number](f float64, p N) float64 {
	x := math.Pow(10, float64(p))
	return math.Round(f*x) / x
}
