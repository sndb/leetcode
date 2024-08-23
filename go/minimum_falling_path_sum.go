package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows)
	for r := range rows {
		dp[r] = make([]int, cols)
		for c := range cols {
			if r == 0 {
				dp[r][c] = matrix[r][c]
			} else {
				dp[r][c] = math.MaxInt
			}
		}
	}
	for r := 1; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for k := c - 1; k <= c+1; k++ {
				if k < 0 || k >= cols {
					continue
				}
				dp[r][c] = min(dp[r][c], dp[r-1][k])
			}
			dp[r][c] += matrix[r][c]
		}
	}
	res := math.MaxInt
	for c := 0; c < cols; c++ {
		res = min(res, dp[rows-1][c])
	}
	return res
}
