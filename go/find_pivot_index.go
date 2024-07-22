package main

func pivotIndex(nums []int) int {
	l, r := 0, 0
	for _, n := range nums {
		r += n
	}
	for i, n := range nums {
		r -= n
		if l == r {
			return i
		}
		l += n
	}
	return -1
}
