package main

func getMaximumXor(nums []int, maximumBit int) []int {
	x := 0
	for _, n := range nums {
		x ^= n
	}
	r := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		r = append(r, (1<<maximumBit-1)^x)
		x ^= nums[i]
	}
	return r
}
