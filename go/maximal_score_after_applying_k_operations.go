package main

import "container/heap"

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func maxKelements(nums []int, k int) int64 {
	h := (*intHeap)(&nums)
	heap.Init(h)

	r := int64(0)
	for range k {
		v := heap.Pop(h).(int)
		r += int64(v)
		x := v / 3
		if v%3 > 0 {
			x += 1
		}
		heap.Push(h, x)
	}
	return r
}
