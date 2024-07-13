package main

import (
	"fmt"
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	indexes := make([]int, len(positions))
	for i := range positions {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return positions[indexes[i]] < positions[indexes[j]]
	})
	stack := []int{}
	for _, i := range indexes {
		if directions[i] == 'R' {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && healths[i] > 0 {
				j := stack[len(stack)-1]
				if healths[j] < healths[i] {
					stack = stack[:len(stack)-1]
					healths[i]--
					healths[j] = 0
				} else if healths[j] > healths[i] {
					healths[i] = 0
					healths[j]--
				} else {
					stack = stack[:len(stack)-1]
					healths[i] = 0
					healths[j] = 0
				}
			}
		}
	}
	result := []int{}
	for _, h := range healths {
		if h > 0 {
			result = append(result, h)
		}
	}
	return result
}

func main() {
	fmt.Println(survivedRobotsHealths([]int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL"))
}
