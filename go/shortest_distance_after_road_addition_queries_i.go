package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	neighbors := map[int][]int{}
	for i := 1; i < n; i++ {
		neighbors[i-1] = []int{i}
	}
	dist := func(curr int) int {
		level := 1
		seen := make([]bool, n)
		queue, next := []int{curr}, []int{}
		for {
			for _, i := range queue {
				for _, neigh := range neighbors[i] {
					if seen[neigh] {
						continue
					}
					if neigh == n-1 {
						return level
					}
					seen[neigh] = true
					next = append(next, neigh)
				}
			}
			queue, next = next, []int{}
			level++
		}
	}
	answer := make([]int, 0, len(queries))
	for _, q := range queries {
		neighbors[q[0]] = append(neighbors[q[0]], q[1])
		answer = append(answer, dist(0))
	}
	return answer
}
