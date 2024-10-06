package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Fields(sentence1)
	s2 := strings.Fields(sentence2)
	n1 := len(s1)
	n2 := len(s2)

	plen := 0
	i1, i2 := 0, 0
	for i1 < n1 && i2 < n2 && s1[i1] == s2[i2] {
		i1++
		i2++
		plen++
	}

	slen := 0
	j1, j2 := n1-1, n2-1
	for j1 >= 0 && j2 >= 0 && s1[j1] == s2[j2] {
		j1--
		j2--
		slen++
	}

	return min(n1, n2) <= plen+slen
}
