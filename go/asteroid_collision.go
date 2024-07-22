package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	q := []int{}
	for _, a := range asteroids {
		for {
			n := len(q)
			if n == 0 || q[n-1] < 0 || a > 0 {
				q = append(q, a)
				break
			}
			if q[n-1] == -a {
				q = q[:n-1]
				break
			} else if q[n-1] < -a {
				q = q[:n-1]
			} else {
				break
			}
		}
	}
	return q
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}
