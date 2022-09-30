package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		if x, ok := m[v]; ok && i-x <= k {
			return true
		}
		m[v] = i
	}
	return false
}
