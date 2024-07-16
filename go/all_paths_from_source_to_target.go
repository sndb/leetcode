package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var walk func([]int)
	walk = func(path []int) {
		curr := path[len(path)-1]
		if curr == len(graph)-1 {
			newp := make([]int, len(path))
			copy(newp, path)
			result = append(result, newp)
		}
		for _, neigh := range graph[curr] {
			walk(append(path, neigh))
		}
	}
	walk([]int{0})
	return result
}
