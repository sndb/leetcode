package main

import "fmt"

func minSteps(n int) int {
	const inf = 1000
	var f func(int, int) int
	f = func(present, copied int) int {
		if present == n {
			return 0
		}
		if present > n {
			return inf
		}
		copy := inf
		if present > copied {
			copy = f(present, present)
		}
		paste := inf
		if copied > 0 {
			paste = f(present+copied, copied)
		}
		return min(copy, paste) + 1
	}
	return f(1, 0)
}

func main() {
	fmt.Println(minSteps(3))  // 3
	fmt.Println(minSteps(1))  // 0
	fmt.Println(minSteps(6))  // 5
	fmt.Println(minSteps(7))  // 7
	fmt.Println(minSteps(11)) // 11
}
