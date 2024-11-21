package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]byte, m)
	for r := range m {
		grid[r] = make([]byte, n)
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = 1
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 1
	}
	for _, g := range guards {
		or, oc := g[0], g[1]
		for r := or - 1; r >= 0; r-- {
			if grid[r][oc] == 1 {
				break
			}
			grid[r][oc] = 2
		}
		for r := or + 1; r < m; r++ {
			if grid[r][oc] == 1 {
				break
			}
			grid[r][oc] = 2
		}
		for c := oc - 1; c >= 0; c-- {
			if grid[or][c] == 1 {
				break
			}
			grid[or][c] = 2
		}
		for c := oc + 1; c < n; c++ {
			if grid[or][c] == 1 {
				break
			}
			grid[or][c] = 2
		}
	}

	x := 0
	for r := range m {
		for c := range n {
			if grid[r][c] == 0 {
				x++
			}
		}
	}
	return x
}
