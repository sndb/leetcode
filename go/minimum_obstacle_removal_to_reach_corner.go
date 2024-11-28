package main

import "container/list"

func minimumObstacles(grid [][]int) int {
	type item struct {
		obstacles int
		row, col  int
	}

	rows := len(grid)
	cols := len(grid[0])
	seen := make([][]bool, rows)
	for r := range rows {
		seen[r] = make([]bool, cols)
	}

	queue := list.New()
	queue.PushFront(item{0, 0, 0})
	for {
		x := queue.Remove(queue.Front()).(item)
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if x.row == rows-1 && x.col == cols-1 {
				return x.obstacles
			}

			nr := x.row + dir[0]
			nc := x.col + dir[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || seen[nr][nc] {
				continue
			}
			seen[nr][nc] = true

			if grid[nr][nc] == 0 {
				queue.PushFront(item{x.obstacles, nr, nc})
			} else {
				queue.PushBack(item{x.obstacles + 1, nr, nc})
			}
		}
	}
}
