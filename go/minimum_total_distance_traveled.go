package main

import "sort"

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	var positions []int
	for _, f := range factory {
		for range f[1] {
			positions = append(positions, f[0])
		}
	}
	sort.Ints(robot)
	sort.Ints(positions)

	const inf = 1e12
	cache := map[[2]int]int64{}
	var calc func(int, int) int64
	calc = func(r, f int) int64 {
		k := [2]int{r, f}
		if v, ok := cache[k]; ok {
			return v
		}

		if r == len(robot) {
			return 0
		}

		if f == len(positions) {
			return inf
		}

		skip := calc(r, f+1)
		assign := int64(abs(robot[r]-positions[f])) + calc(r+1, f+1)
		cache[k] = min(assign, skip)
		return cache[k]
	}
	return calc(0, 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
