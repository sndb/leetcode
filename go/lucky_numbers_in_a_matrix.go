package main

func luckyNumbers(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	minc := make([]int, rows)
	for r := 0; r < rows; r++ {
		i := 0
		for c := 1; c < cols; c++ {
			if matrix[r][c] < matrix[r][i] {
				i = c
			}
		}
		minc[r] = i
	}
	for c := 0; c < cols; c++ {
		i := 0
		for r := 1; r < rows; r++ {
			if matrix[r][c] > matrix[i][c] {
				i = r
			}
		}
		if c == minc[i] {
			return []int{matrix[i][c]}
		}
	}
	return nil
}
