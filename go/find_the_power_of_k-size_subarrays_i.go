package main

func resultsArray(nums []int, k int) []int {
	results := make([]int, len(nums)-k+1)
outer:
	for i := 0; i < len(nums)-k+1; i++ {
		for j := i + 1; j < i+k; j++ {
			if nums[j]-nums[j-1] != 1 {
				results[i] = -1
				continue outer
			}
		}
		results[i] = nums[i+k-1]
	}
	return results
}
