package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	ret := math.MaxInt
	bits := make([]int, 32)
	start := 0
	for i, n := range nums {
		updateBits(bits, n, 1)
		for start <= i && bitsToNumber(bits) >= k {
			updateBits(bits, nums[start], -1)
			ret = min(ret, i-start+1)
			start++
		}
	}
	if ret == math.MaxInt {
		return -1
	}
	return ret
}

func updateBits(b []int, n, v int) {
	for i := range b {
		if n>>i&1 > 0 {
			b[i] += v
		}
	}
}

func bitsToNumber(b []int) int {
	r := 0
	for i, n := range b {
		if n > 0 {
			r |= 1 << i
		}
	}
	return r
}
