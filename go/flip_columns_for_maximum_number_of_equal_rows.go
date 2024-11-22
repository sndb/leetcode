package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	h := map[string]int{}
	for _, r := range matrix {
		s := make([]byte, len(r))
		for i, c := range r {
			s[i] = byte(c ^ r[0])
		}
		h[string(s)]++
	}
	x := 0
	for _, n := range h {
		x = max(x, n)
	}
	return x
}
