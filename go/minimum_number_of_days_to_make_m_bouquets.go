package main

import (
	"fmt"
	"slices"
)

func minDays(bloomDay []int, m, k int) int {
	enough := func(day int) bool {
		a, n := 0, 0
		for _, d := range bloomDay {
			if d <= day {
				a++
			} else {
				a = 0
			}
			if a == k {
				a = 0
				n++
			}
			if n == m {
				return true
			}
		}
		return false
	}
	if len(bloomDay) < m*k {
		return -1
	}
	min := slices.Min(bloomDay)
	max := slices.Max(bloomDay)
	for min < max {
		day := (min + max) / 2
		if enough(day) {
			max = day
		} else {
			min = day + 1
		}
	}
	return min
}

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 2))
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
}
