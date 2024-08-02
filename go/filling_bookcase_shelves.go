package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	m := make([]int, len(books)+1)
	for i := range m {
		m[i] = math.MaxInt
	}
	m[len(books)] = 0
	var f func(int) int
	f = func(i int) int {
		if v := m[i]; v != math.MaxInt {
			return v
		}
		sh, sw := 0, 0
		for j := i; j < len(books); j++ {
			bw := books[j][0]
			bh := books[j][1]
			if sw+bw > shelfWidth {
				break
			}
			sw += bw
			sh = max(sh, bh)
			m[i] = min(m[i], sh+f(j+1))
		}
		return m[i]
	}
	return f(0)
}
