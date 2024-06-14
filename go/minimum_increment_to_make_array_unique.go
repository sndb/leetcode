package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]
		if prev >= curr {
			nums[i] = prev + 1
			moves += prev + 1 - curr
		}
	}
	return moves
}
