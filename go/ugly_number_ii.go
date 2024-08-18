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

func nthUglyNumber(n int) int {
	num := 1
	seen := map[int]bool{}
	nums := &intHeap{}
	for range n - 1 {
		for _, p := range []int{2, 3, 5} {
			if x := num * p; !seen[x] {
				heap.Push(nums, x)
				seen[x] = true
			}
		}
		num = heap.Pop(nums).(int)
	}
	return num
}
