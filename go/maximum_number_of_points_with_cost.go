package main

import "fmt"

func maxPoints(points [][]int) int64 {
	rows := len(points)
	cols := len(points[0])
	lmax := make([]int64, cols)
	rmax := make([]int64, cols)
	prev := make([]int64, cols)
	for c := 0; c < cols; c++ {
		prev[c] = int64(points[0][c])
	}
	for r := 1; r < rows; r++ {
		lmax[0] = prev[0]
		rmax[cols-1] = prev[cols-1]
		for c := 1; c < cols; c++ {
			lmax[c] = max(lmax[c-1]-1, prev[c])
		}
		for c := cols - 2; c >= 0; c-- {
			rmax[c] = max(rmax[c+1]-1, prev[c])
		}
		for c := 0; c < cols; c++ {
			prev[c] = int64(points[r][c]) + max(lmax[c], rmax[c])
		}
	}
	best := prev[0]
	for c := 1; c < cols; c++ {
		best = max(best, prev[c])
	}
	return best
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}))
	fmt.Println(maxPoints([][]int{{1, 5}, {2, 3}, {4, 2}}))
}
