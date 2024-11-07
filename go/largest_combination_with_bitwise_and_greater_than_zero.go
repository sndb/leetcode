package main

func largestCombination(candidates []int) int {
	ret := 0
	bits := make([]int, 24)
	for _, n := range candidates {
		for i := range bits {
			bits[i] += n & 1
			ret = max(ret, bits[i])
			n >>= 1
		}
	}
	return ret
}
