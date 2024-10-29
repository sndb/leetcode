package main

func maxMoves(grid [][]int) int {
	moves := 0
	seen := map[[3]int]bool{}
	var walk func(int, int, int)
	walk = func(prev, r, c int) {
		if r < 0 || r >= len(grid) || c >= len(grid[0]) || prev >= grid[r][c] {
			return
		}

		k := [3]int{prev, r, c}
		if seen[k] {
			return
		}
		seen[k] = true

		moves = max(moves, c)
		walk(grid[r][c], r-1, c+1)
		walk(grid[r][c], r, c+1)
		walk(grid[r][c], r+1, c+1)
	}
	for r := 0; r < len(grid); r++ {
		walk(0, r, 0)
	}
	return moves
}
