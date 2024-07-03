package main

import (
	"fmt"
	"sort"
)

func minDifference(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}
	sort.Ints(nums)
	return min(
		nums[n-4]-nums[0],
		nums[n-3]-nums[1],
		nums[n-2]-nums[2],
		nums[n-1]-nums[3],
	)
}

func main() {
	fmt.Println(minDifference([]int{6, 6, 0, 1, 1, 4, 6}))   // 2
	fmt.Println(minDifference([]int{9, 48, 92, 48, 81, 31})) // 17
}
