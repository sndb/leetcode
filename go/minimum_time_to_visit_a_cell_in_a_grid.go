package main

import "container/heap"

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	rows := len(grid)
	cols := len(grid[0])
	seen := make([][]bool, rows)
	for r := range rows {
		seen[r] = make([]bool, cols)
	}

	queue := items{{0, 0, 0}}
	for {
		x := heap.Pop(&queue).(item)
		if x.row == rows-1 && x.col == cols-1 {
			return x.time
		}

		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr := x.row + dir[0]
			nc := x.col + dir[1]
			if nr < 0 || nr == rows || nc < 0 || nc == cols || seen[nr][nc] {
				continue
			}

			nt1 := x.time + 1
			nt2 := grid[nr][nc] + (grid[nr][nc]-x.time+1)%2
			heap.Push(&queue, item{max(nt1, nt2), nr, nc})
			seen[nr][nc] = true
		}
	}
}

type item struct {
	time int
	row  int
	col  int
}

type items []item

func (h items) Len() int           { return len(h) }
func (h items) Less(i, j int) bool { return h[i].time < h[j].time }
func (h items) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *items) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *items) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
