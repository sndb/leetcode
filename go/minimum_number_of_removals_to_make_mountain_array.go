package main

func minimumMountainRemovals(nums []int) int {
	lis := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	lds := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		lds[i] = 1
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[i] {
				lds[i] = max(lds[i], lds[j]+1)
			}
		}
	}

	r := 0
	for i := 0; i < len(nums); i++ {
		if lis[i] > 1 && lds[i] > 1 {
			r = max(r, lis[i]+lds[i]-1)
		}
	}
	return len(nums) - r
}
