package main

import "slices"

func arrayRankTransform(arr []int) []int {
	rank := 1
	ranks := map[int]int{}
	for _, n := range slices.Sorted(slices.Values(arr)) {
		if _, ok := ranks[n]; !ok {
			ranks[n] = rank
			rank++
		}
	}
	result := make([]int, len(arr))
	for i, n := range arr {
		result[i] = ranks[n]
	}
	return result
}
