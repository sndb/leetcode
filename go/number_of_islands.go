package main

func explore(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	explore(grid, r+1, c)
	explore(grid, r-1, c)
	explore(grid, r, c+1)
	explore(grid, r, c-1)
}

func numIslands(grid [][]byte) int {
	num := 0
	for r, row := range grid {
		for c, val := range row {
			if val == '1' {
				explore(grid, r, c)
				num++
			}
		}
	}
	return num
}
