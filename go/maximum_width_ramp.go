package main

func maxWidthRamp(nums []int) int {
	dec := []int{0}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[dec[len(dec)-1]] {
			dec = append(dec, i)
		}
	}

	ret := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for len(dec) > 0 && nums[dec[len(dec)-1]] <= nums[i] {
			ret = max(ret, i-dec[len(dec)-1])
			dec = dec[:len(dec)-1]
		}
	}
	return ret
}
