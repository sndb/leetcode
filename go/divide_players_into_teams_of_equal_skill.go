package main

import "slices"

func dividePlayers(skill []int) int64 {
	slices.Sort(skill)
	n := len(skill)

	target := skill[0] + skill[n-1]
	result := int64(0)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		s1 := skill[i]
		s2 := skill[j]
		if s1+s2 != target {
			return -1
		}
		result += int64(s1 * s2)
	}
	return result
}
