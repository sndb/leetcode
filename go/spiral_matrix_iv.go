package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	mat := make([][]int, m)
	for i := range m {
		mat[i] = make([]int, n)
		for j := range n {
			mat[i][j] = -1
		}
	}
	r, c := 0, 0
	dr, dc := 0, 1
	for head != nil {
		mat[r][c] = head.Val
		nr, nc := r+dr, c+dc
		if nr < 0 || nr >= m || nc < 0 || nc >= n || mat[nr][nc] != -1 {
			dr, dc = dc, -dr
			nr, nc = r+dr, c+dc
		}
		r, c = nr, nc
		head = head.Next
	}
	return mat
}
