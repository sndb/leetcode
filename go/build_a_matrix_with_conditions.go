package main

import "fmt"

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rows := topoSort(rowConditions, k)
	if len(rows) != k {
		return nil
	}
	cols := topoSort(colConditions, k)
	if len(cols) != k {
		return nil
	}
	icol := map[int]int{}
	for c, n := range cols {
		icol[n] = c
	}
	matrix := make([][]int, k)
	for i := range k {
		matrix[i] = make([]int, k)
	}
	for r, n := range rows {
		matrix[r][icol[n]] = n
	}
	return matrix
}

func topoSort(conds [][]int, k int) []int {
	adj := make([][]int, k+1)
	in := make([]int, k+1)
	for _, c := range conds {
		adj[c[0]] = append(adj[c[0]], c[1])
		in[c[1]]++
	}
	queue := []int{}
	for n := 1; n <= k; n++ {
		if in[n] == 0 {
			queue = append(queue, n)
		}
	}
	list := []int{}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		list = append(list, n)
		for _, m := range adj[n] {
			in[m]--
			if in[m] == 0 {
				queue = append(queue, m)
			}
		}
	}
	return list
}

func main() {
	fmt.Println(buildMatrix(3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}))
	fmt.Println(buildMatrix(3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}))
	fmt.Println(buildMatrix(8, [][]int{{1, 2}, {7, 3}, {4, 3}, {5, 8}, {7, 8}, {8, 2}, {5, 8}, {3, 2}, {1, 3}, {7, 6}, {4, 3}, {7, 4}, {4, 8}, {7, 3}, {7, 5}}, [][]int{{5, 7}, {2, 7}, {4, 3}, {6, 7}, {4, 3}, {2, 3}, {6, 2}}))
}
