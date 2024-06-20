package main

import (
	"fmt"
	"sort"
)

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	canPlace := func(force int) bool {
		count := m - 1
		prev := position[0]
		for _, pos := range position[1:] {
			if pos-prev >= force {
				prev = pos
				count--
			}
			if count == 0 {
				return true
			}
		}
		return false
	}

	min := 1
	max := position[len(position)-1] - position[0]
	for min <= max {
		guess := (min + max) / 2
		if canPlace(guess) {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}
	return max
}

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
	fmt.Println(maxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
}
