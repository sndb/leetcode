package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	graph := map[int][]int{}
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[to] = append(graph[to], from)
	}

	ancestors := make([][]int, n)
	for v := range n {
		anc := []int{}
		seen := make([]bool, n)
		stack := []int{v}
		for len(stack) > 0 {
			n := len(stack)
			u := stack[n-1]
			stack = stack[:n-1]
			for _, w := range graph[u] {
				if seen[w] {
					continue
				}
				seen[w] = true
				stack = append(stack, w)
				anc = append(anc, w)
			}
		}
		sort.Ints(anc)
		ancestors[v] = anc
	}
	return ancestors
}
