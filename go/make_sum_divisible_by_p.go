package main

func minSubarray(nums []int, p int) int {
	total := 0
	for _, n := range nums {
		total = (total + n) % p
	}
	if total == 0 {
		return 0
	}

	acc := 0
	result := len(nums)
	index := map[int]int{0: -1}
	for i, n := range nums {
		acc = (acc + n) % p
		comp := (acc - total + p) % p
		if j, ok := index[comp]; ok {
			result = min(result, i-j)
		}
		index[acc] = i
	}
	if result == len(nums) {
		return -1
	}
	return result
}
