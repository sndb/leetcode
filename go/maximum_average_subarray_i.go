package main

func findMaxAverage(nums []int, k int) float64 {
	s := 0
	for i := 0; i < k; i++ {
		s += nums[i]
	}
	t := s
	for i := 0; i+k < len(nums); i++ {
		s -= nums[i]
		s += nums[i+k]
		t = max(s, t)
	}
	return float64(t) / float64(k)
}
