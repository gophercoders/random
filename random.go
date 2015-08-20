// Copyright 2015 Kulawe Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package random providesa simple way to generate randon number, both
// positive and negative within a range.
//
// The functions provided by this module are not crytographically safe as
// the package uses the pseudo-random number generator from the math.rand
// package. The random number generator is seeded with the time in nanoseconds
// that this package is first initialised.
package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomNumberInRange returns an int in the range min to max (inclusive).
// In other words the values of both min and max are included in the range and
// can be returned by the function. The function can return negative numbers if
// min is less than zero, so a range between -12 and 1 can return the numbers
// -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1.
//
// The function will panic is max < min or if max == min.
func GetRandomNumberInRange(min, max int) int {
	var n int
	var r int
	// max < min is bad
	if max < min {
		panic("max < min")
	}
	//max == min is also bad
	if max == min {
		panic("max == min")
	}
	// looking for a number in the range zero to max (where max is positive) were
	// both zero amd max are included in the range.
	if min == 0 && max > 0 {
		n = rand.Intn(max + 1) // gives range of 0..max
		return n
	}
	// looking for a number in the range min (were min is negative) to zero were
	// both min and zero are in the range
	if max == 0 && min < 0 {
		n = rand.Intn((min - 1) * -1) // gives range 0..abs(min)+1. rand.Intn requires a postive number
		return n * -1                 // the range consists of only or negative numbers so negate the answer
	}
	if min < 0 {
		if max < 0 {
			// looking for a number in the range min to max were both numbers are negative
			// and both numbers are included in the range.
			r = (max - min) - max // can be negative
			if r < 0 {            // if it is negative make it positive
				r = r * -1
			}
		} else {
			// looking for a number in the range min to max were min is negative
			// and max is positive and both numbers are included in the range.
			r = (max - min) + max // can never be negative
		}
	} else if min > 0 {
		// looking for a number in the range min to max where both min and max
		// are positive and both numbers are included in the range.
		r = (max - min) + min
	}
	n = rand.Intn(r) // gives a number in the range 0..r-1
	return (n + min) // gives a number in the range n..r-1+min. When min is negative min is subtracted not added.
}
