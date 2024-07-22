package main

func largestAltitude(gain []int) int {
	a := 0
	r := 0
	for _, g := range gain {
		a += g
		r = max(r, a)
	}
	return r
}
