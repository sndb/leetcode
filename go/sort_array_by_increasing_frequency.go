package main

import "sort"

func frequencySort(nums []int) []int {
	f := map[int]int{}
	for _, n := range nums {
		f[n]++
	}
	sort.Slice(nums, func(i, j int) bool {
		m, n := nums[i], nums[j]
		return f[m] < f[n] || f[m] == f[n] && m > n
	})
	return nums
}
