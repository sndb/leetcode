package main

import (
	"sort"
)

func frequencySort(s string) string {
	freq := make([]int, 1<<8)
	for _, v := range s {
		freq[v]++
	}
	r := []byte(s)
	sort.Slice(r, func(i, j int) bool {
		if freq[r[i]] == freq[r[j]] {
			return r[i] > r[j]
		}
		return freq[r[i]] > freq[r[j]]
	})
	return string(r)
}
