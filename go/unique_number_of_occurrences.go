package main

func uniqueOccurrences(arr []int) bool {
	h := map[int]int{}
	for _, x := range arr {
		h[x]++
	}
	o := map[int]bool{}
	for _, n := range h {
		o[n] = true
	}
	return len(o) == len(h)
}
