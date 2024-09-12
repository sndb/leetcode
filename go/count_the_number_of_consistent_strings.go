package main

func countConsistentStrings(allowed string, words []string) int {
	good := [26]bool{}
	for _, c := range allowed {
		good[c-'a'] = true
	}
	n := 0
outer:
	for _, w := range words {
		for _, c := range w {
			if !good[c-'a'] {
				continue outer
			}
		}
		n++
	}
	return n
}
