package main

import "fmt"

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	nr, nc := len(rowSum), len(colSum)
	matrix := make([][]int, nr)
	for i := 0; i < nr; i++ {
		matrix[i] = make([]int, nc)
	}
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			v := min(rowSum[r], colSum[c])
			matrix[r][c] = v
			rowSum[r] -= v
			colSum[c] -= v
		}
	}
	return matrix
}

func main() {
	fmt.Println(restoreMatrix([]int{3, 8}, []int{4, 7}))
}
