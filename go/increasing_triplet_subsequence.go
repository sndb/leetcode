package main

import "math"

func increasingTriplet(nums []int) bool {
	l, m := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n <= l {
			l = n
		} else if n <= m {
			m = n
		} else {
			return true
		}
	}
	return false
}
