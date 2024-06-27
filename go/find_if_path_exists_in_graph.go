package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := map[int][]int{}
	for _, e := range edges {
		v1, v2 := e[0], e[1]
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}

	seen := map[int]bool{}
	queue := []int{source}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == destination {
			return true
		}
		for _, u := range graph[v] {
			if seen[u] {
				continue
			}
			seen[u] = true
			queue = append(queue, u)
		}
	}
	return false
}
