package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	acc := 0
	curr := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		acc += gas[i] - cost[i]
		curr += gas[i] - cost[i]
		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}
	if acc < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             // -1
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3
	fmt.Println(canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))       // 3
}
