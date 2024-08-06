package main

import "sort"

func minimumPushes(word string) int {
	s := make([]int, 26)
	for _, c := range word {
		s[c-'a']++
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	r := 0
	for i, n := range s {
		r += (i/8 + 1) * n
	}
	return r
}
