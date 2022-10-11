// Package number contains generic functions
// primarily for use with numbers
package number

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Number is any number type
type Number interface {
	constraints.Integer | constraints.Float
}

// Round rounds a float to the specified digits of precision
func Round[N Number](f float64, p N) float64 {
	x := math.Pow(10, float64(p))
	return math.Round(f*x) / x
}

// Max returns the maximum in a list of numbers
func Max[N constraints.Ordered](nums ...N) (n N) {
	if len(nums) <= 0 {
		return n
	}
	n = nums[0]
	for _, num := range nums {
		if num > n {
			n = num
		}
	}
	return n
}

// Min returns the minimum in a list of numbers
func Min[N constraints.Ordered](nums ...N) (n N) {
	if len(nums) <= 0 {
		return n
	}
	n = nums[0]
	for _, num := range nums {
		if num < n {
			n = num
		}
	}
	return n
}

// Sum returns the sum of a list of numbers
func Sum[N constraints.Ordered](nums ...N) (n N) {
	for _, num := range nums {
		n += num
	}
	return n
}
