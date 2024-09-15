package main

func findTheLongestSubstring(s string) int {
	t := ['z' + 1]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
	h := [32]int{0: -1}
	for i := 1; i < 32; i++ {
		h[i] = -2
	}
	x := 0
	l := 0
	for i, c := range s {
		x ^= t[c]
		if j := h[x]; j == -2 {
			h[x] = i
		} else {
			l = max(l, i-j)
		}
	}
	return l
}
