package main

func minAddToMakeValid(s string) int {
	add := 0
	cl, op := 0, 0
	for _, c := range s {
		if c == '(' {
			op++
		} else {
			cl++
		}
		if cl > op {
			op++
			add++
		}
	}
	return add + op - cl
}
