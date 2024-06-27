package main

func findCenter(edges [][]int) int {
	seen := map[int]bool{}
	for _, e := range edges {
		v1 := e[0]
		if seen[v1] {
			return v1
		}
		v2 := e[1]
		if seen[v2] {
			return v2
		}
		seen[v1] = true
		seen[v2] = true
	}
	return 0
}
