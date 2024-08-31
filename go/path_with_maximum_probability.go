package main

import "container/heap"

type maxHeap []node

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *maxHeap) Pop() any {
	x := (*h)[0]
	*h = (*h)[1:]
	return x
}

type node struct {
	index int
	prob  float64
}

func maxProbability(n int, edges [][]int, succ []float64, start, end int) float64 {
	adj := map[int][]node{}
	for i, e := range edges {
		u, v, p := e[0], e[1], succ[i]
		adj[u] = append(adj[u], node{v, p})
		adj[v] = append(adj[v], node{u, p})
	}

	prob := make([]float64, n)
	prob[start] = 1.0

	seen := make([]bool, n)
	queue := &maxHeap{node{index: start, prob: 1}}
	for queue.Len() > 0 {
		u := heap.Pop(queue).(node)
		if seen[u.index] {
			continue
		}
		seen[u.index] = true

		for _, v := range adj[u.index] {
			prob[v.index] = max(prob[v.index], prob[u.index]*v.prob)
			heap.Push(queue, node{v.index, prob[v.index]})
		}
	}
	return prob[end]
}
