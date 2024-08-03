package main

import "slices"

func canBeEqual(target []int, arr []int) bool {
	s1 := make([]int, 1001)
	for _, x := range target {
		s1[x]++
	}
	s2 := make([]int, 1001)
	for _, x := range arr {
		s2[x]++
	}
	return slices.Equal(s1, s2)
}
