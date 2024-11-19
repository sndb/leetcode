package main

func maximumSubarraySum(nums []int, k int) int64 {
	count := map[int]int{}
	currSum := int64(0)
	maxSum := int64(0)
	for i := 0; i < k; i++ {
		count[nums[i]]++
		currSum += int64(nums[i])
	}
	if len(count) == k {
		maxSum = currSum
	}
	for i := k; i < len(nums); i++ {
		count[nums[i-k]]--
		count[nums[i]]++
		if count[nums[i-k]] == 0 {
			delete(count, nums[i-k])
		}
		currSum -= int64(nums[i-k])
		currSum += int64(nums[i])
		if len(count) == k {
			maxSum = max(maxSum, currSum)
		}
	}
	return maxSum
}
