package main

import "sort"

func unjumble(t []int, n int) int {
	if n == 0 {
		return t[0]
	}
	r := 0
	k := 1
	for n > 0 {
		r += k * t[n%10]
		n /= 10
		k *= 10
	}
	return r
}

func sortJumbled(mapping []int, nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return unjumble(mapping, nums[i]) < unjumble(mapping, nums[j])
	})
	return nums
}
