package main

import "container/heap"

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	h *intHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{&intHeap{}, k}
	for _, n := range nums {
		s.Add(n)
	}
	return s
}

func (s *KthLargest) Add(val int) int {
	heap.Push(s.h, val)
	if s.h.Len() > s.k {
		heap.Pop(s.h)
	}
	return (*s.h)[0]
}
