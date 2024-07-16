package main

import "fmt"

func pivot(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	p := pivot(nums)
	if target >= nums[p] && target <= nums[r] {
		l = p
	} else {
		r = p - 1
	}
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(pivot([]int{4, 5, 6, 7, 0, 1, 2}))
}
