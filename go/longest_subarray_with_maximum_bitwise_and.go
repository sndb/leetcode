package main

func longestSubarray(nums []int) int {
	var m, c, l int
	for _, n := range nums {
		if n > m {
			m = n
			c = 1
			l = 1
		} else if n == m {
			c++
			l = max(l, c)
		} else {
			c = 0
		}
	}
	return l
}
