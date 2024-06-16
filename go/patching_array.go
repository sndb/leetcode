package main

func minPatches(nums []int, n int) int {
	i, x, p := 0, 1, 0
	for x <= n {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x += x
			p++
		}
	}
	return p
}
