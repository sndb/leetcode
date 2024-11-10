package main

func minEnd(n int, x int) int64 {
	r := int64(x)
	for n > 1 {
		r = (r + 1) | int64(x)
		n--
	}
	return r
}
