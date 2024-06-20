package main

import "slices"

func minimizedMaximum(n int, quantities []int) int {
	distribute := func(x int) bool {
		c := 0
		for _, q := range quantities {
			c += (q + x - 1) / x
		}
		return c <= n
	}
	min := 1
	max := slices.Max(quantities)
	for min <= max {
		avg := (min + max) / 2
		if distribute(avg) {
			max = avg - 1
		} else {
			min = avg + 1
		}
	}
	return min
}
