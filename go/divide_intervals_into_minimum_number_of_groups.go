package main

import "slices"

func minGroups(intervals [][]int) int {
	lefts := make([]int, len(intervals))
	rights := make([]int, len(intervals))
	for i, v := range intervals {
		lefts[i] = v[0]
		rights[i] = v[1]
	}
	slices.Sort(lefts)
	slices.Sort(rights)

	i, j, n, r := 0, 0, 0, 0
	for i < len(lefts) && j < len(rights) {
		if lefts[i] <= rights[j] {
			i++
			n++
			r = max(r, n)
		} else {
			j++
			n--
		}
	}
	return r
}
