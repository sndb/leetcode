package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	const maxT = 31
	const maxV = 51
	m := [maxT]map[[maxV]int]bool{}
	for i := range m {
		m[i] = map[[maxV]int]bool{}
	}
	for _, c := range candidates {
		if c > target {
			continue
		}
		a := [maxT][][maxV]int{}
		for s1, l := range m {
			s2 := c + s1
			if s2 > target {
				continue
			}
			for x := range l {
				x[c]++
				a[s2] = append(a[s2], x)
			}
		}
		for s, l := range a {
			for _, x := range l {
				m[s][x] = true
			}
		}
		x := [maxV]int{}
		x[c] = 1
		m[c][x] = true
	}
	r := [][]int{}
	for x := range m[target] {
		a := []int{}
		for n, c := range x {
			for range c {
				a = append(a, n)
			}
		}
		r = append(r, a)
	}
	return r
}

func main() {
	fmt.Println(combinationSum2([]int{1, 2, 3, 2, 2}, 6))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
