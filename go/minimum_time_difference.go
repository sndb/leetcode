package main

import (
	"fmt"
	"slices"
)

func findMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))
	for i, tp := range timePoints {
		var h, m int
		fmt.Sscanf(tp, "%d:%d", &h, &m)
		minutes[i] = h*60 + m
	}
	slices.Sort(minutes)

	diff := 24 * 60
	prev := minutes[len(minutes)-1] - diff
	for _, curr := range minutes {
		diff = min(diff, curr-prev)
		prev = curr
	}
	return diff
}
