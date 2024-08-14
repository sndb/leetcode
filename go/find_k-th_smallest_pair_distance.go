package main

import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	l, h := 0, nums[len(nums)-1]-nums[0]
	for l < h {
		m := (l + h) / 2
		c := countPairs(nums, m)
		if c < k {
			l = m + 1
		} else {
			h = m
		}
	}
	return l
}

func countPairs(nums []int, d int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] > d {
				break
			}
			r++
		}
	}
	return r
}

func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 1))
	fmt.Println(smallestDistancePair([]int{1, 1, 1}, 2))
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
}
