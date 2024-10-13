package main

import "math"

func smallestRange(nums [][]int) []int {
	rmin, rmax := 0, math.MaxInt
	for {
		imin, imax := 0, 0
		for i := 1; i < len(nums); i++ {
			if nums[i][0] < nums[imin][0] {
				imin = i
			}
			if nums[i][0] > nums[imax][0] {
				imax = i
			}
		}
		if nums[imax][0]-nums[imin][0] < rmax-rmin {
			rmin, rmax = nums[imin][0], nums[imax][0]
		}
		nums[imin] = nums[imin][1:]
		if len(nums[imin]) == 0 {
			break
		}
	}
	return []int{rmin, rmax}
}
