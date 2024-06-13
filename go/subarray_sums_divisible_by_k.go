package main

import "fmt"

func subarraysDivByK(nums []int, k int) int {
	seen := map[int]int{0: 1}
	m := 0
	r := 0
	for _, n := range nums {
		m = (m + n) % k
		if m < 0 {
			m += k
		}
		r += seen[m]
		seen[m]++
	}
	return r
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{5}, 9))                  // 0
	fmt.Println(subarraysDivByK([]int{8, 9, 7, 8, 9}, 8))      // 7
}
