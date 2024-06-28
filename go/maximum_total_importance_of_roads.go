package main

import (
	"cmp"
	"slices"
)

func maximumImportance(n int, roads [][]int) int64 {
	count := map[int]int{}
	for _, r := range roads {
		count[r[0]]++
		count[r[1]]++
	}

	cities := []int{}
	for c := range count {
		cities = append(cities, c)
	}
	slices.SortFunc(cities, func(a, b int) int {
		return cmp.Compare(count[b], count[a])
	})

	var r int64 = 0
	for i, c := range cities {
		r += int64((n - i) * count[c])
	}
	return r
}
