package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	wc1 := map[string]int{}
	for _, w := range strings.Fields(s1) {
		wc1[w]++
	}
	wc2 := map[string]int{}
	for _, w := range strings.Fields(s2) {
		wc2[w]++
	}
	words := []string{}
	for w, c := range wc1 {
		if c == 1 && wc2[w] == 0 {
			words = append(words, w)
		}
	}
	for w, c := range wc2 {
		if c == 1 && wc1[w] == 0 {
			words = append(words, w)
		}
	}
	return words
}
