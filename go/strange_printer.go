package main

import "slices"

func strangePrinter(s string) int {
	chars := slices.Compact([]byte(s))
	n := len(chars)

	cache := make([][]int, n)
	for i := range n {
		cache[i] = make([]int, n)
	}

	var minTurns func(int, int) int
	minTurns = func(start, end int) int {
		if start > end {
			return 0
		}
		res := cache[start][end]
		if res != 0 {
			return res
		}
		res = 1 + minTurns(start+1, end)
		for i := start + 1; i <= end; i++ {
			if chars[start] == chars[i] {
				res = min(res, minTurns(start, i-1)+minTurns(i+1, end))
			}
		}
		cache[start][end] = res
		return res
	}
	return minTurns(0, n-1)
}
