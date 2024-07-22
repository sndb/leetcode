package main

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	o := 0
	l := 0
	r := len(nums) - 1
	for l < r {
		s := nums[l] + nums[r]
		if s > k {
			r--
		} else if s < k {
			l++
		} else {
			o++
			l++
			r--
		}
	}
	return o
}
