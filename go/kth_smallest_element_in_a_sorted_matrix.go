package main

import (
	"container/heap"
	"fmt"
)

type entry struct{ r, c, v int }
type queue []entry

func (h queue) Len() int           { return len(h) }
func (h queue) Less(i, j int) bool { return h[i].v < h[j].v }
func (h queue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *queue) Push(x any) {
	*h = append(*h, x.(entry))
}

func (h *queue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	h := &queue{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, entry{i, 0, matrix[i][0]})
	}
	for {
		e := heap.Pop(h).(entry)
		if k--; k == 0 {
			return e.v
		}
		if nc := e.c + 1; nc < len(matrix) {
			heap.Push(h, entry{e.r, nc, matrix[e.r][nc]})
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(matrix, 8))
}
