package main

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows := len(grid1)
	cols := len(grid1[0])
	var walk func(int, int) bool
	walk = func(r, c int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid2[r][c] == 0 {
			return true
		}
		grid2[r][c] = 0
		r1 := walk(r-1, c)
		r2 := walk(r+1, c)
		r3 := walk(r, c-1)
		r4 := walk(r, c+1)
		return grid1[r][c] == 1 && r1 && r2 && r3 && r4
	}
	islands := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid2[r][c] == 1 && walk(r, c) {
				islands++
			}
		}
	}
	return islands
}
