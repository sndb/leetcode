package main

import "slices"

func heightChecker(heights []int) int {
	n := len(heights)
	expected := make([]int, n)
	copy(expected, heights)
	slices.Sort(expected)

	r := 0
	for i := range n {
		if heights[i] != expected[i] {
			r++
		}
	}
	return r
}
