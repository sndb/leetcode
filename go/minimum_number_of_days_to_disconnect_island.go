package main

func minDays(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	if !isConnected(grid) {
		return 0
	}
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				grid[r][c] = 0
				if !isConnected(grid) {
					return 1
				}
				grid[r][c] = 1
			}
		}
	}
	return 2
}

func isConnected(grid [][]int) bool {
	rows := len(grid)
	cols := len(grid[0])
	size := 0
	for r := range rows {
		for c := range cols {
			size += grid[r][c]
		}
	}
	seen := map[[2]int]bool{}
	var walk func([2]int)
	walk = func(p [2]int) {
		if !seen[p] && p[0] >= 0 && p[0] < rows && p[1] >= 0 && p[1] < cols && grid[p[0]][p[1]] == 1 {
			seen[p] = true
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				walk([2]int{p[0] + d[0], p[1] + d[1]})
			}
		}
	}
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				walk([2]int{r, c})
				return len(seen) == size
			}
		}
	}
	return false
}
