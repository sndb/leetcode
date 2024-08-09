package main

func numMagicSquaresInside(grid [][]int) int {
	n := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if isMagicSquare(grid, r, c) {
				n++
			}
		}
	}
	return n
}

func isMagicSquare(grid [][]int, r, c int) bool {
	if r+2 >= len(grid) || c+2 >= len(grid[0]) {
		return false
	}
	want := 0
	for i := 0; i < 3; i++ {
		want += grid[r][c+i]
	}
	seen := [10]bool{}
	for i := 0; i < 3; i++ {
		s := 0
		for j := 0; j < 3; j++ {
			x := grid[r+i][c+j]
			if x < 1 || x > 9 || seen[x] {
				return false
			}
			seen[x] = true
			s += x
		}
		if s != want {
			return false
		}
	}
	for j := 0; j < 3; j++ {
		s := 0
		for i := 0; i < 3; i++ {
			s += grid[r+i][c+j]
		}
		if s != want {
			return false
		}
	}
	d1 := 0
	d2 := 0
	for i := 0; i < 3; i++ {
		d1 += grid[r+i][c+i]
		d2 += grid[r+i][c+2-i]
	}
	if d1 != want || d2 != want {
		return false
	}
	return true
}
