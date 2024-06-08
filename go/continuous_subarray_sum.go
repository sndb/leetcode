package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	mod := 0
	seen := map[int]int{0: -1}
	for i, n := range nums {
		mod = (mod + n) % k
		j, ok := seen[mod]
		if ok && i-j > 1 {
			return true
		}
		if !ok {
			seen[mod] = i
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))  // true
	fmt.Println(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13)) // false
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))  // true
	fmt.Println(checkSubarraySum([]int{1, 0}, 2))            // false
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))      // true
	fmt.Println(checkSubarraySum([]int{0, 0}, 1))            // true
	fmt.Println(checkSubarraySum([]int{1, 2, 12}, 6))        // false
	fmt.Println(checkSubarraySum([]int{5, 2, 4}, 5))         // false
}
