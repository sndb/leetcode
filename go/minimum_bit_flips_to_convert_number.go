package main

func minBitFlips(start int, goal int) int {
	n := 0
	for start > 0 || goal > 0 {
		n += start&1 ^ goal&1
		start >>= 1
		goal >>= 1
	}
	return n
}
