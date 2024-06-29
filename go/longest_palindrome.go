package main

func longestPalindrome(s string) int {
	h := map[rune]int{}
	for _, c := range s {
		h[c]++
	}
	l, m := 0, 0
	for _, n := range h {
		r := n % 2
		l += n - r
		m = max(m, r)
	}
	return l + m
}
