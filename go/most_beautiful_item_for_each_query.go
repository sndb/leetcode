package main

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	items = append(items, []int{0, 0})
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	for i := 1; i < len(items); i++ {
		items[i][1] = max(items[i-1][1], items[i][1])
	}
	answer := make([]int, len(queries))
	for i, q := range queries {
		low, high := 0, len(items)
		for low < high {
			mid := (low + high) / 2
			if items[mid][0] > q {
				high = mid
			} else {
				low = mid + 1
			}
		}
		answer[i] = items[high-1][1]
	}
	return answer
}
