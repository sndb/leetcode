package main

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	for p, i := 1, 0; i < len(nums); i++ {
		out[i] = p
		p *= nums[i]
	}
	for p, i := 1, len(nums)-1; i >= 0; i-- {
		out[i] *= p
		p *= nums[i]
	}
	return out
}
