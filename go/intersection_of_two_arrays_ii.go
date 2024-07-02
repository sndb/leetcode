package main

func intersect(nums1 []int, nums2 []int) []int {
	h := map[int]int{}
	for _, n := range nums1 {
		h[n]++
	}
	r := []int{}
	for _, n := range nums2 {
		if h[n] > 0 {
			h[n]--
			r = append(r, n)
		}
	}
	return r
}
