package main

import "strconv"

func maximumSwap(num int) int {
	r := num
	a := []byte(strconv.Itoa(num))
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			n, _ := strconv.Atoi(string(a))
			a[i], a[j] = a[j], a[i]
			r = max(r, n)
		}
	}
	return r
}
