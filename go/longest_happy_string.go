package main

import "sort"

func longestDiverseString(a, b, c int) string {
	pp, p := byte(3), byte(3)
	chars := []byte{0, 1, 2}
	count := []int{a, b, c}
	r := []byte{}

outer:
	for {
		sort.Slice(chars, func(i, j int) bool {
			return count[chars[i]] > count[chars[j]]
		})
		for _, c := range chars {
			if (p != c || pp != c) && count[c] > 0 {
				r = append(r, 'a'+c)
				pp, p = p, c
				count[c]--
				continue outer
			}
		}
		break
	}
	return string(r)
}
