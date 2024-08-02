package main

func minSwaps(nums []int) int {
	total1 := 0
	for _, n := range nums {
		if n == 1 {
			total1++
		}
	}
	curr1 := 0
	for _, n := range nums[:total1] {
		if n == 1 {
			curr1++
		}
	}
	max1 := curr1
	for i, j := 0, total1; i < len(nums); i, j = i+1, j+1 {
		curr1 -= nums[i]
		curr1 += nums[j%len(nums)]
		max1 = max(max1, curr1)
	}
	return total1 - max1
}
