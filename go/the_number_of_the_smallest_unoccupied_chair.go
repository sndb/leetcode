package main

import "slices"

func smallestChair(times [][]int, targetFriend int) int {
	targetArrival := times[targetFriend][0]
	slices.SortFunc(times, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return +1
		}
		return 0
	})
	leavings := make([]int, len(times))
	for _, t := range times {
		a, l := t[0], t[1]
		for i, ll := range leavings {
			if ll <= a {
				if a == targetArrival {
					return i
				}
				leavings[i] = l
				break
			}
		}
	}
	return -1
}
