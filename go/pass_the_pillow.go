package main

import "fmt"

func passThePillow(n int, time int) int {
	m, d := 1, -1
	for range time {
		if m == 1 || m == n {
			d = -d
		}
		m += d
	}
	return m
}

func main() {
	fmt.Println(passThePillow(4, 5))   // 2
	fmt.Println(passThePillow(3, 2))   // 3
	fmt.Println(passThePillow(18, 38)) // 5
	fmt.Println(passThePillow(2, 341)) // 2
}
