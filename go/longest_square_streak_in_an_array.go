package main

import "slices"

func longestSquareStreak(nums []int) int {
	h := map[int]bool{}
	for _, n := range nums {
		h[n] = true
	}

	r := -1
	slices.Sort(nums)
	for _, n := range nums {
		c := 0
		for h[n] {
			n *= n
			c++
		}
		if c > 1 {
			r = max(r, c)
		}
	}
	return r
}
