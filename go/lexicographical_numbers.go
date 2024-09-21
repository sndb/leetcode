package main

import "fmt"

func lexicalOrder(n int) []int {
	var result []int
	var walk func(int)
	walk = func(x int) {
		if x <= n {
			result = append(result, x)
			for i := 0; i <= 9; i++ {
				walk(10*x + i)
			}
		}
	}
	for i := 1; i <= 9; i++ {
		walk(i)
	}
	return result
}

func main() {
	fmt.Println(lexicalOrder(25))
}
