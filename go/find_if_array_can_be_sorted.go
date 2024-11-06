package main

import "math/bits"

type segment struct {
	setBits  int
	largest  int
	smallest int
}

func canSortArray(nums []int) bool {
	segments := []segment{}
	for _, x := range nums {
		n := len(segments)
		c := bits.OnesCount(uint(x))
		if n == 0 || segments[n-1].setBits != c {
			segments = append(segments, segment{c, x, x})
		} else {
			segments[n-1].largest = max(segments[n-1].largest, x)
			segments[n-1].smallest = min(segments[n-1].smallest, x)
		}
	}
	for i := 1; i < len(segments); i++ {
		if segments[i-1].largest > segments[i].smallest {
			return false
		}
	}
	return true
}
