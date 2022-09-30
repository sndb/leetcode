package main

func intersection(nums1 []int, nums2 []int) []int {
	seen := map[int]int{}
	for _, v := range nums1 {
		seen[v] = 1
	}
	for _, v := range nums2 {
		if seen[v] == 1 {
			seen[v] += 1
		}
	}
	r := []int{}
	for i, v := range seen {
		if v == 2 {
			r = append(r, i)
		}
	}
	return r
}
