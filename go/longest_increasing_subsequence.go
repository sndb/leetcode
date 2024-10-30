package main

func lengthOfLIS(nums []int) int {
	r := 0
	l := make([]int, len(nums))
	for i, m := range nums {
		l[i] = 1
		for j, n := range nums[:i] {
			if n < m {
				l[i] = max(l[i], l[j]+1)
			}
		}
		r = max(r, l[i])
	}
	return r
}
