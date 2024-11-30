package main

import "slices"

func validArrangement(pairs [][]int) [][]int {
	adj := map[int][]int{}
	count := map[int]int{}

	for _, p := range pairs {
		adj[p[0]] = append(adj[p[0]], p[1])
		count[p[0]]++
		count[p[1]]--
	}

	start := pairs[0][0]
	for u, c := range count {
		if c == 1 {
			start = u
			break
		}
	}

	ret := [][]int{}
	var walk func(int)
	walk = func(u int) {
		for len(adj[u]) > 0 {
			n := len(adj[u])
			v := adj[u][n-1]
			adj[u] = adj[u][:n-1]
			walk(v)
			ret = append(ret, []int{u, v})
		}
	}
	walk(start)
	slices.Reverse(ret)
	return ret
}
