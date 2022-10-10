package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	s := map[int]int{}
	for _, n := range nums1 {
		for _, m := range nums2 {
			s[n+m]++
		}
	}
	i := 0
	for _, n := range nums3 {
		for _, m := range nums4 {
			i += s[-n-m]
		}
	}
	return i
}

func main() {
	nums1 := []int{-1, -1}
	nums2 := []int{-1, 1}
	nums3 := []int{-1, 1}
	nums4 := []int{1, -1}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
