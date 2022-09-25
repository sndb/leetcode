package main

import "fmt"

// 1st implementation: map cache, recursion
func rob1(nums []int) int {
	cache := map[int]int{}
	return rob1c(nums, cache)
}

func rob1c(nums []int, cache map[int]int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	r, ok := cache[len(nums)]
	if ok {
		return r
	}
	cache[len(nums)] = max(rob1c(nums[1:], cache), nums[0]+rob1c(nums[2:], cache))
	return cache[len(nums)]
}

// 2nd implementation: slice cache, iteration
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	cache := make([]int, len(nums))
	cache[0] = nums[0]
	cache[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cache[i] = max(cache[i-1], cache[i-2]+nums[i])
	}
	return cache[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	inputs := [][]int{
		{2, 7, 9, 3, 1},    // 12
		{1, 2, 3, 1},       // 4
		{3, 1, 1, 3, 5},    // 9
		{3, 1, 1, 3, 5, 2}, // 9
		{3, 1, 3, 1},       // 6
		{3, 1, 1, 3},       // 6
	}
	for _, v := range inputs {
		fmt.Println(rob1(v))
	}
}
