package main

import (
	"container/heap"
	"math"
)

func shortestSubarray(nums []int, k int) int {
	queue := pairHeap{}
	ret := math.MaxInt
	acc := 0
	for i, n := range nums {
		acc += n
		if acc >= k {
			ret = min(ret, i+1)
		}
		for len(queue) > 0 && acc-queue[0].sum >= k {
			ret = min(ret, i-heap.Pop(&queue).(pair).index)
		}
		heap.Push(&queue, pair{acc, i})
	}
	if ret == math.MaxInt {
		return -1
	}
	return ret
}

type pair struct {
	sum   int
	index int
}

type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
