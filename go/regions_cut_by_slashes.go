package main

const (
	top = iota
	right
	left
	bottom
	dirs
)

func regionsBySlashes(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	state := make([][][4]int, rows)
	for r := range rows {
		state[r] = make([][4]int, cols)
	}
	reg := 0
	var walk func(int, int, int)
	walk = func(r, c, d int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || state[r][c][d] != 0 {
			return
		}
		state[r][c][d] = reg
		typ := grid[r][c]
		switch d {
		case top:
			walk(r-1, c, bottom)
			if typ == '/' {
				walk(r, c, left)
			} else if typ == '\\' {
				walk(r, c, right)
			} else {
				walk(r, c, left)
				walk(r, c, right)
				walk(r, c, bottom)
			}
		case right:
			walk(r, c+1, left)
			if typ == '/' {
				walk(r, c, bottom)
			} else if typ == '\\' {
				walk(r, c, top)
			} else {
				walk(r, c, bottom)
				walk(r, c, top)
				walk(r, c, left)
			}
		case left:
			walk(r, c-1, right)
			if typ == '/' {
				walk(r, c, top)
			} else if typ == '\\' {
				walk(r, c, bottom)
			} else {
				walk(r, c, top)
				walk(r, c, bottom)
				walk(r, c, right)
			}
		case bottom:
			walk(r+1, c, top)
			if typ == '/' {
				walk(r, c, right)
			} else if typ == '\\' {
				walk(r, c, left)
			} else {
				walk(r, c, right)
				walk(r, c, left)
				walk(r, c, top)
			}
		}
	}
	for r := range rows {
		for c := range cols {
			for d := range dirs {
				if state[r][c][d] == 0 {
					reg++
					walk(r, c, d)
				}
			}
		}
	}
	return reg
}
