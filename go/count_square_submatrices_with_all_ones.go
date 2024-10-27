package main

func countSquares(matrix [][]int) int {
	s := 0
	m := len(matrix)
	n := len(matrix[0])
	for r := range m {
		for c := range n {
			for k := 1; r+k <= m && c+k <= n; k++ {
				if allOnes(matrix, r, c, k) {
					s++
				}
			}
		}
	}
	return s
}

func allOnes(matrix [][]int, r, c, k int) bool {
	for i := r; i < r+k; i++ {
		for j := c; j < c+k; j++ {
			if matrix[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
