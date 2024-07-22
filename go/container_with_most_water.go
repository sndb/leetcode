package main

func maxArea(height []int) int {
	a := 0
	l := 0
	r := len(height) - 1
	for l < r {
		if height[l] > height[r] {
			a = max(a, (r-l)*height[r])
			r--
		} else {
			a = max(a, (r-l)*height[l])
			l++
		}
	}
	return a
}
