package main

import "sort"

type Pair struct{ A, B []int }

func (p *Pair) Len() int           { return len(p.A) }
func (p *Pair) Less(i, j int) bool { return p.A[i] > p.A[j] }
func (p *Pair) Swap(i, j int)      { p.A[i], p.A[j], p.B[i], p.B[j] = p.A[j], p.A[i], p.B[j], p.B[i] }

func maxProfitAssignment(difficulty, profit, worker []int) int {
	sort.Sort(&Pair{profit, difficulty})
	sort.Slice(worker, func(i, j int) bool { return worker[i] > worker[j] })
	r, i, j := 0, 0, 0
	for i < len(profit) && j < len(worker) {
		if difficulty[i] > worker[j] {
			i++
		} else {
			r += profit[i]
			j++
		}
	}
	return r
}
