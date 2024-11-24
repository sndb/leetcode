package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	neg, sum, low := 0, 0, math.MaxInt
	for _, r := range matrix {
		for _, x := range r {
			if x < 0 {
				neg++
				sum -= x
				low = min(low, -x)
			} else {
				sum += x
				low = min(low, x)
			}
		}
	}
	return int64(sum - neg%2*low*2)
}
