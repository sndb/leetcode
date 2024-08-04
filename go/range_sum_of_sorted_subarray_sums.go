package main

import "slices"

func rangeSum(nums []int, n int, left int, right int) int {
	a := []int{}
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			a = append(a, s)
		}
	}
	slices.Sort(a)
	r := 0
	for i := left - 1; i < right; i++ {
		r = (r + a[i]) % 1000000007
	}
	return r
}
