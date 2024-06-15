package main

import (
	"container/heap"
	"sort"
)

type Pair struct{ A, B []int }

func (p Pair) Len() int           { return len(p.A) }
func (p Pair) Less(i, j int) bool { return p.B[i] < p.B[j] }
func (p Pair) Swap(i, j int)      { p.A[i], p.A[j], p.B[i], p.B[j] = p.A[j], p.A[i], p.B[j], p.B[i] }

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	a := *h
	n := len(a)
	x := a[n-1]
	*h = a[0 : n-1]
	return x
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	sort.Sort(Pair{profits, capital})
	i, r, h := 0, w, &MaxHeap{}
	for k > 0 {
		for i < len(capital) && capital[i] <= w {
			heap.Push(h, profits[i])
			i++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
		if w > r {
			r = w
			k--
		}
	}
	return r
}
