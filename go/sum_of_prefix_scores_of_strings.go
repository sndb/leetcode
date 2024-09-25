package main

func sumPrefixScores(words []string) []int {
	count := map[string]int{}
	for _, w := range words {
		for n := 1; n <= len(w); n++ {
			count[w[:n]]++
		}
	}
	ret := make([]int, len(words))
	for i, w := range words {
		for n := 1; n <= len(w); n++ {
			ret[i] += count[w[:n]]
		}
	}
	return ret
}
