package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return lowerBound(nums, upper+1) - lowerBound(nums, lower)
}

func lowerBound(nums []int, val int) int64 {
	low := 0
	high := len(nums) - 1
	count := int64(0)
	for low < high {
		sum := nums[low] + nums[high]
		if sum < val {
			count += int64(high - low)
			low++
		} else {
			high--
		}
	}
	return count
}
