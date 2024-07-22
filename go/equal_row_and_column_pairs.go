package main

const m = 200

func equalPairs(grid [][]int) int {
	n := len(grid)
	rows := map[[m]int]int{}
	for r := 0; r < n; r++ {
		k := [m]int{}
		for c := 0; c < n; c++ {
			k[c] = grid[r][c]
		}
		rows[k]++
	}
	p := 0
	for c := 0; c < n; c++ {
		k := [m]int{}
		for r := 0; r < n; r++ {
			k[r] = grid[r][c]
		}
		p += rows[k]
	}
	return p
}
