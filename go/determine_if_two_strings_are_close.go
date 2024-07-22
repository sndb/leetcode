package main

import "slices"

func stringSignature(s string) (int, [26]int) {
	a := 0
	c := [26]int{}
	for _, r := range s {
		k := r - 'a'
		a |= 1 << k
		c[k]++
	}
	slices.Sort(c[:])
	return a, c
}

func closeStrings(s1 string, s2 string) bool {
	a1, c1 := stringSignature(s1)
	a2, c2 := stringSignature(s2)
	return a1 == a2 && c1 == c2
}
