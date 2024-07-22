package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	h1 := map[int]bool{}
	for _, n := range nums1 {
		h1[n] = true
	}
	h2 := map[int]bool{}
	for _, n := range nums2 {
		h2[n] = true
	}
	d := make([][]int, 2)
	for n := range h1 {
		if !h2[n] {
			d[0] = append(d[0], n)
		}
	}
	for n := range h2 {
		if !h1[n] {
			d[1] = append(d[1], n)
		}
	}
	return d
}
